#!/usr/bin/env python3
"""Build, run, score, and stop a private native modelpicker server."""

import argparse
import csv
from datetime import datetime, timezone
import hashlib
from http.client import HTTPException
import io
import json
import math
import os
from pathlib import Path
import platform
import secrets
import shlex
import signal
import socket
import statistics
import subprocess
import sys
import tempfile
import time
from urllib import error, request


ROOT = Path(__file__).resolve().parents[1]
MODEL_FILES = ("laya.onnx", "laya.onnx.data", "tokenizer/tokenizer.json")


def sha256(path):
    with path.open("rb") as stream:
        return hashlib.file_digest(stream, "sha256").hexdigest()


def native_paths():
    system = {"Darwin": "darwin", "Linux": "linux"}.get(platform.system())
    arch = {"arm64": "arm64", "aarch64": "arm64", "x86_64": "amd64"}.get(platform.machine())
    if (system, arch) not in {("darwin", "arm64"), ("linux", "amd64"), ("linux", "arm64")}:
        raise ValueError("Native evals support macOS ARM64 and Linux AMD64/ARM64")
    directory = ROOT / "lib" / f"{system}-{arch}"
    runtime = directory / "lib" / ("libonnxruntime.dylib" if system == "darwin" else "libonnxruntime.so")
    return system, arch, directory, runtime


def stop_process(process, timeout=35):
    try:
        os.killpg(process.pid, signal.SIGTERM)
    except ProcessLookupError:
        process.wait()
        return False
    forced = False
    try:
        process.wait(timeout=timeout)
    except subprocess.TimeoutExpired:
        forced = True
    # A finished leader can leave descendants in its process group.
    try:
        os.killpg(process.pid, signal.SIGKILL)
        forced = True
    except ProcessLookupError:
        pass
    process.wait()
    return forced


def run_command(command, log, env):
    print("Running " + shlex.join(command), flush=True)
    process = subprocess.Popen(command, cwd=ROOT, env=env, stdout=log, stderr=subprocess.STDOUT, start_new_session=True)
    try:
        if process.wait() != 0:
            raise RuntimeError(f"Command failed: {shlex.join(command)}; see build.log")
    finally:
        stop_process(process)


def build_server(work_dir, model_dir, log):
    system, arch, libraries, _ = native_paths()
    env = os.environ.copy()
    env.update(GOOS=system, GOARCH=arch, CGO_ENABLED="1")
    run_command(["make", "deps"], log, env)
    if not all((model_dir / name).is_file() for name in MODEL_FILES):
        run_command(["make", "model", f"MODEL_DIR={model_dir}"], log, env)
    env["CGO_LDFLAGS"] = " ".join(filter(None, [env.get("CGO_LDFLAGS"), shlex.quote(f"-L{libraries}")]))
    binary = work_dir / "modelpicker"
    run_command(["go", "build", "-o", str(binary), "."], log, env)
    return binary


def load_inputs(cases_path, config_path):
    cases_bytes, config_bytes = cases_path.read_bytes(), config_path.read_bytes()
    policy = json.loads(config_bytes)
    if not isinstance(policy, dict) or not isinstance(policy.get("tiers"), list):
        raise ValueError("The eval config must be JSON with a tiers array")
    if policy.get("default_tier"):
        raise ValueError("Remove default_tier from the eval config so inference failures stay visible")
    tiers = {}
    for tier in policy["tiers"]:
        if not isinstance(tier, dict) or not isinstance(tier.get("name"), str) or not tier["name"].strip():
            raise ValueError("Each tier needs a nonempty name")
        models = tier.get("models")
        if tier["name"] in tiers or not isinstance(models, list) or not models or not all(isinstance(m, str) and m.strip() for m in models):
            raise ValueError("Tier names must be unique and model lists must contain nonempty strings")
        tiers[tier["name"]] = models
    if not tiers:
        raise ValueError("The eval config needs at least one tier")
    reader = csv.DictReader(io.StringIO(cases_bytes.decode("utf-8-sig"), newline=""))
    required = {"question_id", "category", "prompt", "expected_tier", "notes"}
    if not required.issubset(reader.fieldnames or []) or len(set(reader.fieldnames)) != len(reader.fieldnames):
        raise ValueError("CSV needs unique question_id, category, prompt, expected_tier, and notes columns")
    cases, seen = [], set()
    for row in reader:
        if None in row or any(value is None for value in row.values()):
            raise ValueError(f"Malformed CSV record ending on line {reader.line_num}")
        if not row["question_id"].strip() or row["question_id"] in seen or not row["prompt"].strip():
            raise ValueError(f"CSV line {reader.line_num}: IDs must be unique and IDs/prompts must not be empty")
        if row["expected_tier"] not in tiers:
            raise ValueError(f"CSV line {reader.line_num}: expected_tier must be one of {', '.join(tiers)}")
        seen.add(row["question_id"])
        cases.append(row)
    if not cases:
        raise ValueError("CSV contains no cases")
    return cases, tiers, cases_bytes, config_bytes


def pick(opener, url, key, prompt, tiers, timeout):
    body = json.dumps({"prompt": prompt}, ensure_ascii=False).encode()
    req = request.Request(url + "/v1/pick", data=body, headers={"Authorization": f"Bearer {key}", "Content-Type": "application/json"})
    with opener.open(req, timeout=timeout) as response:
        try:
            data = json.load(response)
        except HTTPException as exc:
            raise ValueError(f"Incomplete HTTP response: {exc}") from exc
    if not isinstance(data, dict) or set(data) != {"tier", "model"}:
        raise ValueError("Response must contain exactly tier and model")
    if not isinstance(data["tier"], str) or data["tier"] not in tiers:
        raise ValueError("Response contains an unknown tier")
    if not isinstance(data["model"], str) or data["model"] not in tiers[data["tier"]]:
        raise ValueError("Response model is not a member of its tier")
    return data


def wait_ready(process, opener, url, timeout):
    deadline = time.monotonic() + timeout
    while True:
        remaining = deadline - time.monotonic()
        if remaining <= 0:
            break
        if process.poll() is not None:
            raise RuntimeError(f"Server exited with code {process.returncode} before readiness; see server.log")
        try:
            with opener.open(url + "/healthz", timeout=min(1, remaining)) as response:
                if response.status == 200:
                    return
        except (OSError, error.URLError, HTTPException):
            pass
        time.sleep(0.1)
    raise RuntimeError(f"Server was not ready within {timeout:g}s; see server.log")


def evaluate_case(case, opener, url, key, tiers, timeout):
    result = {**case, "predicted_tier": None, "model": None, "correct": False, "http_status": None, "error": None}
    started = time.perf_counter()
    try:
        data = pick(opener, url, key, case["prompt"], tiers, timeout)
        result.update(predicted_tier=data["tier"], model=data["model"], http_status=200, correct=data["tier"] == case["expected_tier"])
    except error.HTTPError as exc:
        result.update(http_status=exc.code, error=f"HTTP {exc.code}: {exc.read().decode(errors='replace').strip()}")
    except ValueError as exc:
        result.update(http_status=200, error=str(exc))
    except (OSError, error.URLError, HTTPException) as exc:
        result["error"] = str(exc)
    result["latency_ms"] = (time.perf_counter() - started) * 1000
    return result


def metrics(results, tiers):
    names = list(tiers)
    confusion = {name: dict.fromkeys(names, 0) for name in names}
    errors_by_tier = dict.fromkeys(names, 0)
    groups, per_tier = {}, {}
    for result in results:
        if result["error"] is not None:
            errors_by_tier[result["expected_tier"]] += 1
        else:
            confusion[result["expected_tier"]][result["predicted_tier"]] += 1
        suite = "advanced" if result["question_id"].startswith(("frontiercs/", "researchrubrics/")) else "baseline"
        groups.setdefault(suite, []).append(result)
    for name in names:
        correct = confusion[name][name]
        total = sum(confusion[name].values()) + errors_by_tier[name]
        selected = sum(confusion[expected][name] for expected in names)
        precision = correct / selected if selected else 0
        recall = correct / total if total else 0
        per_tier[name] = {"count": total, "precision": precision, "recall": recall, "f1": 2 * precision * recall / (precision + recall) if precision + recall else 0}
    latencies = sorted(result["latency_ms"] for result in results)
    return {
        "accuracy": sum(result["correct"] for result in results) / len(results),
        "errors": sum(result["error"] is not None for result in results),
        "macro_f1": statistics.mean(tier["f1"] for tier in per_tier.values()),
        "per_tier": per_tier, "confusion_matrix": confusion, "errors_by_tier": errors_by_tier,
        "by_suite": {name: {"count": len(group), "accuracy": sum(r["correct"] for r in group) / len(group)} for name, group in groups.items()},
        "latency_ms": {"p50": statistics.median(latencies), "p95": latencies[math.ceil(0.95 * len(latencies)) - 1]},
    }


def save_summary(output, state):
    temporary = output / "summary.json.tmp"
    temporary.write_text(json.dumps(state, indent=2) + "\n")
    temporary.replace(output / "summary.json")


def run_evaluation(args):
    cases, tiers, cases_bytes, config_bytes = load_inputs(args.cases, args.config)
    output = args.output or ROOT / "evals" / "runs" / (datetime.now(timezone.utc).strftime("%Y%m%dT%H%M%SZ-") + secrets.token_hex(4))
    output = output.resolve()
    output.mkdir(parents=True, mode=0o700)
    (output / "cases.csv").write_bytes(cases_bytes)
    (output / "config.json").write_bytes(config_bytes)
    print(f"Run artifacts: {output}", flush=True)
    state = {
        "status": "running", "phase": "setup", "started_at": datetime.now(timezone.utc).isoformat(),
        "total_cases": len(cases), "completed_cases": 0, "metrics": None,
        "cases_sha256": hashlib.sha256(cases_bytes).hexdigest(),
        "config_sha256": hashlib.sha256(config_bytes).hexdigest(),
        "runner_sha256": sha256(Path(__file__)), "model_dir": str(args.model_dir),
    }
    results, process = [], None
    save_summary(output, state)
    try:
        with tempfile.TemporaryDirectory(prefix="modelpicker-eval-") as temporary, \
                (output / "build.log").open("w") as build_log, (output / "server.log").open("w") as server_log:
            work_dir = Path(temporary)
            binary = build_server(work_dir, args.model_dir, build_log)
            _, _, _, runtime = native_paths()
            state["binary_sha256"] = sha256(binary)
            state["model_sha256"] = {name: sha256(args.model_dir / name) for name in MODEL_FILES}
            state["runtime_sha256"] = sha256(runtime) if runtime.is_file() else None
            key = secrets.token_hex(32)
            with socket.socket() as reservation:
                reservation.bind(("127.0.0.1", 0))
                port = reservation.getsockname()[1]
            url = f"http://127.0.0.1:{port}"
            env = os.environ.copy()
            env.update(API_KEY=key, CONFIG_PATH=str(output / "config.json"), MODEL_DIR=str(args.model_dir),
                       ONNX_RUNTIME_PATH=str(runtime), LISTEN_ADDR=f"127.0.0.1:{port}", WRITE_TIMEOUT=f"{args.request_timeout:g}s")
            state.update(phase="starting", url=url)
            save_summary(output, state)
            started = time.monotonic()
            try:
                process = subprocess.Popen([str(binary)], cwd=work_dir, env=env, stdout=server_log, stderr=subprocess.STDOUT, start_new_session=True)
                state["server_pid"] = process.pid
                save_summary(output, state)
                opener = request.build_opener(request.ProxyHandler({}))
                wait_ready(process, opener, url, args.startup_timeout)
                pick(opener, url, key, "Hello.", tiers, args.request_timeout)
                state.update(phase="evaluating", startup_seconds=time.monotonic() - started)
                save_summary(output, state)
                with (output / "results.jsonl").open("w") as stream:
                    for case in cases:
                        if process.poll() is not None:
                            raise RuntimeError(f"Server exited with code {process.returncode}; see server.log")
                        result = evaluate_case(case, opener, url, key, tiers, args.request_timeout)
                        results.append(result)
                        stream.write(json.dumps(result, ensure_ascii=False) + "\n")
                        stream.flush()
                        state["completed_cases"] = len(results)
                        save_summary(output, state)
                failed = sum(result["error"] is not None for result in results)
                state["status"] = "failed" if failed else "completed"
                if failed:
                    state["error"] = f"{failed} request(s) failed"
            finally:
                if process is not None:
                    state["forced_shutdown"] = stop_process(process)
                    state["server_returncode"] = process.returncode
                    if state["status"] == "completed" and (process.returncode != 0 or state["forced_shutdown"]):
                        state.update(status="failed", error=f"Server did not shut down cleanly (code {process.returncode}); see server.log")
    except KeyboardInterrupt:
        state.update(status="interrupted", error="Run interrupted")
    except (OSError, ValueError, RuntimeError, HTTPException) as exc:
        state.update(status="failed", error=str(exc))
    finally:
        state.update(phase="finished", finished_at=datetime.now(timezone.utc).isoformat())
        if len(results) == len(cases):
            state["metrics"] = metrics(results, tiers)
        save_summary(output, state)
    if state["metrics"] is not None:
        report = state["metrics"]
        print(f"Accuracy: {report['accuracy']:.1%}; macro F1: {report['macro_f1']:.3f}; request errors: {report['errors']}")
        for name, group in report["by_suite"].items():
            print(f"  {name}: {group['accuracy']:.1%} ({group['count']} cases)")
    if state.get("error"):
        print(state["error"], file=sys.stderr)
    return 130 if state["status"] == "interrupted" else 0 if state["status"] == "completed" else 2


def positive_seconds(value):
    seconds = float(value)
    if not math.isfinite(seconds) or seconds <= 0:
        raise argparse.ArgumentTypeError("timeout must be a positive finite number")
    return seconds


def main(argv=None):
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--cases", type=Path, default=ROOT / "evals/suites/mt-bench-baseline.csv")
    parser.add_argument("--config", type=Path, default=ROOT / "evals/config.json", help="JSON config accepted by the server's YAML loader")
    parser.add_argument("--model-dir", type=Path, default=ROOT / "inference/model")
    parser.add_argument("--output", type=Path, help="new artifact directory; must not already exist")
    parser.add_argument("--startup-timeout", type=positive_seconds, default=180)
    parser.add_argument("--request-timeout", type=positive_seconds, default=60)
    args = parser.parse_args(argv)
    args.model_dir = args.model_dir.resolve()
    previous = signal.getsignal(signal.SIGTERM)

    def interrupt(signum, frame):
        raise KeyboardInterrupt

    signal.signal(signal.SIGTERM, interrupt)
    try:
        return run_evaluation(args)
    except (OSError, ValueError) as exc:
        print(str(exc), file=sys.stderr)
        return 2
    except KeyboardInterrupt:
        return 130
    finally:
        signal.signal(signal.SIGTERM, previous)


if __name__ == "__main__":
    sys.exit(main())
