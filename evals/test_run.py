import contextlib
import csv
import io
import json
import os
from pathlib import Path
import signal
import subprocess
import sys
import tempfile
import time
import unittest
from unittest import mock

import run as runner


FAKE_SERVER = r'''#!/usr/bin/env python3
import json
import os
import signal
import sys
from http.server import BaseHTTPRequestHandler, HTTPServer

config = json.load(open(os.environ["CONFIG_PATH"]))
print("pid=" + str(os.getpid()), flush=True)
print("cwd=" + os.getcwd(), flush=True)
mode = config["classifier"]["question"]
if mode == "exit":
    sys.exit(9)
signal.signal(signal.SIGTERM, lambda *_: sys.exit(0))
models = {t["name"]: t["models"][0] for t in config["tiers"]}

class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(503 if mode == "not-ready" else 200)
        self.end_headers()

    def do_POST(self):
        if self.headers.get("Authorization") != "Bearer " + os.environ["API_KEY"]:
            self.send_response(401)
            self.end_headers()
            return
        prompt = json.loads(self.rfile.read(int(self.headers["Content-Length"])))["prompt"]
        if prompt == "http-error":
            self.send_response(500)
            self.end_headers()
            return
        if prompt == "truncated":
            self.send_response(200)
            self.send_header("Content-Length", "100")
            self.end_headers()
            self.wfile.write(b"{")
            self.close_connection = True
            return
        tier = {"easy": "tier_1", "normal": "tier_2"}.get(prompt, "tier_3")
        body = json.dumps({"tier": tier, "model": "invalid" if prompt == "invalid" else models[tier]}).encode()
        self.send_response(200)
        self.send_header("Content-Type", "application/json")
        self.end_headers()
        self.wfile.write(body)

host, port = os.environ["LISTEN_ADDR"].rsplit(":", 1)
with HTTPServer((host, int(port)), Handler) as server:
    server.serve_forever()
'''


class RunnerTest(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory()
        self.addCleanup(self.temp.cleanup)
        self.root = Path(self.temp.name)
        self.output = self.root / "output"
        self.binary = self.root / "fake-picker"
        self.binary.write_text(FAKE_SERVER)
        self.binary.chmod(0o755)
        self.model = self.root / "model"
        for name in ("laya.onnx", "laya.onnx.data", "tokenizer/tokenizer.json"):
            path = self.model / name
            path.parent.mkdir(parents=True, exist_ok=True)
            path.write_text("fixture")
        self.config = self.root / "config.json"
        self.policy = {
            "classifier": {"question": "capability"},
            "tiers": [
                {"name": tier, "description": tier, "models": [f"eval/{tier}"]}
                for tier in ("tier_1", "tier_2", "tier_3")
            ],
        }
        self.config.write_text(json.dumps(self.policy))
        self.cases = self.root / "cases.csv"
        self.write_cases([
            ("1", "general", "easy", "tier_1", ""),
            ("2", "coding", "normal", "tier_3", "intentional mismatch"),
            ("frontiercs/test", "research", "hard", "tier_3", ""),
        ])

    def write_cases(self, rows):
        with self.cases.open("w", newline="") as stream:
            writer = csv.writer(stream)
            writer.writerow(["question_id", "category", "prompt", "expected_tier", "notes"])
            writer.writerows(rows)

    def invoke(self, *extra):
        args = [
            "--cases", str(self.cases), "--config", str(self.config),
            "--model-dir", str(self.model), "--output", str(self.output),
            "--startup-timeout", "3", "--request-timeout", "1", *extra,
        ]
        with mock.patch.object(runner, "build_server", return_value=self.binary), \
                contextlib.redirect_stdout(io.StringIO()), contextlib.redirect_stderr(io.StringIO()):
            return runner.main(args)

    def summary(self):
        return json.loads((self.output / "summary.json").read_text())

    def assert_server_stopped(self):
        log = (self.output / "server.log").read_text()
        pid = self.summary()["server_pid"]
        with self.assertRaises(ProcessLookupError):
            os.kill(pid, 0)
        self.assertNotIn("cwd=" + str(Path(__file__).resolve().parents[1]) + "\n", log)

    def test_real_http_scoring_snapshots_and_cleanup(self):
        with mock.patch.dict(os.environ, {
            "API_KEY": "parent-secret", "CONFIG_PATH": "/invalid/parent/config",
            "LISTEN_ADDR": "not-a-listener", "HTTP_PROXY": "http://127.0.0.1:1",
        }):
            self.assertEqual(self.invoke(), 0)
        summary = self.summary()
        self.assertEqual(summary["status"], "completed")
        self.assertEqual(summary["metrics"]["accuracy"], 2 / 3)
        self.assertEqual(summary["metrics"]["confusion_matrix"]["tier_3"]["tier_2"], 1)
        self.assertEqual(summary["metrics"]["by_suite"]["baseline"]["accuracy"], 0.5)
        self.assertEqual(summary["metrics"]["by_suite"]["advanced"]["accuracy"], 1)
        self.assertEqual((self.output / "cases.csv").read_bytes(), self.cases.read_bytes())
        self.assertEqual((self.output / "config.json").read_bytes(), self.config.read_bytes())
        self.assertEqual(len((self.output / "results.jsonl").read_text().splitlines()), 3)
        self.assertNotIn("parent-secret", json.dumps(summary))
        self.assert_server_stopped()

    def test_http_and_model_membership_failures_are_recorded_and_counted(self):
        self.write_cases([
            ("1", "general", "http-error", "tier_1", ""),
            ("2", "general", "invalid", "tier_3", ""),
            ("3", "general", "hard", "tier_3", ""),
        ])
        self.assertEqual(self.invoke(), 2)
        summary = self.summary()
        self.assertEqual(summary["status"], "failed")
        self.assertEqual(summary["metrics"]["accuracy"], 1 / 3)
        self.assertEqual(summary["metrics"]["errors"], 2)
        results = [json.loads(line) for line in (self.output / "results.jsonl").read_text().splitlines()]
        self.assertEqual(results[0]["http_status"], 500)
        self.assertEqual(results[1]["http_status"], 200)
        self.assertIn("model", results[1]["error"])
        self.assert_server_stopped()

    def test_startup_exit_is_reported(self):
        self.policy["classifier"]["question"] = "exit"
        self.config.write_text(json.dumps(self.policy))
        self.assertEqual(self.invoke(), 2)
        self.assertEqual(self.summary()["status"], "failed")
        self.assertIn("exited", self.summary()["error"])
        self.assert_server_stopped()

    def test_startup_timeout_stops_server(self):
        self.policy["classifier"]["question"] = "not-ready"
        self.config.write_text(json.dumps(self.policy))
        self.assertEqual(self.invoke("--startup-timeout", "0.3"), 2)
        self.assertIn("ready", self.summary()["error"])
        self.assert_server_stopped()

    def test_sigterm_saves_partial_state_and_stops_server(self):
        previous_handler = signal.getsignal(signal.SIGTERM)
        with mock.patch.object(runner, "evaluate_case", side_effect=lambda *args: os.kill(os.getpid(), signal.SIGTERM)):
            self.assertEqual(self.invoke(), 130)
        self.assertEqual(self.summary()["status"], "interrupted")
        self.assertIsNone(self.summary()["metrics"])
        self.assertEqual(signal.getsignal(signal.SIGTERM), previous_handler)
        self.assert_server_stopped()

    def test_invalid_labels_are_rejected_before_build(self):
        self.write_cases([("1", "general", "easy", "", "")])
        with mock.patch.object(runner, "build_server") as build:
            with contextlib.redirect_stderr(io.StringIO()):
                self.assertEqual(runner.main([
                    "--cases", str(self.cases), "--config", str(self.config),
                    "--output", str(self.output),
                ]), 2)
            build.assert_not_called()
        self.assertFalse(self.output.exists())

    def test_truncated_response_is_recorded_and_remaining_cases_run(self):
        self.write_cases([
            ("1", "general", "truncated", "tier_3", ""),
            ("2", "general", "easy", "tier_1", ""),
        ])
        self.assertEqual(self.invoke(), 2)
        self.assertEqual(self.summary()["status"], "failed")
        self.assertEqual(self.summary()["metrics"]["accuracy"], 0.5)
        rows = [json.loads(line) for line in (self.output / "results.jsonl").read_text().splitlines()]
        self.assertEqual(rows[0]["http_status"], 200)
        self.assertIsNotNone(rows[0]["error"])
        self.assertTrue(rows[1]["correct"])
        self.assert_server_stopped()

    def test_error_tier_is_distinct_from_request_failures(self):
        result = {
            "question_id": "1", "expected_tier": "error", "predicted_tier": None,
            "correct": False, "error": "HTTP 500", "latency_ms": 1,
        }
        report = runner.metrics([result], {"error": ["model"]})
        self.assertEqual(report["accuracy"], 0)
        self.assertEqual(report["per_tier"]["error"]["precision"], 0)
        self.assertEqual(report["per_tier"]["error"]["recall"], 0)
        self.assertEqual(report["per_tier"]["error"]["f1"], 0)

    def test_cleanup_stops_descendants_after_their_leader_exits(self):
        child_code = "import signal,time; signal.signal(signal.SIGTERM, signal.SIG_IGN); print('ready',flush=True); time.sleep(60)"
        parent_code = (
            "import subprocess,sys,time; "
            f"p=subprocess.Popen([sys.executable,'-c',{child_code!r}],stdout=subprocess.PIPE,text=True); "
            "p.stdout.readline(); print(p.pid,flush=True); time.sleep(60)"
        )
        process = subprocess.Popen([sys.executable, "-c", parent_code], stdout=subprocess.PIPE, text=True, start_new_session=True)
        try:
            child_pid = int(process.stdout.readline())
            forced = runner.stop_process(process, timeout=0.2)
            deadline = time.monotonic() + 2
            while time.monotonic() < deadline:
                state = subprocess.run(["ps", "-o", "stat=", "-p", str(child_pid)], capture_output=True, text=True).stdout.strip()
                if not state or state.startswith("Z"):
                    break
                time.sleep(0.02)
            else:
                self.fail("Owned descendant survived process-group cleanup")
            self.assertTrue(forced)
        finally:
            try:
                os.killpg(process.pid, signal.SIGKILL)
            except ProcessLookupError:
                pass
            process.wait()
            process.stdout.close()


if __name__ == "__main__":
    unittest.main()
