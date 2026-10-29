# Tests and CI

Requires the [native toolchain](../README.md#native-setup). From the repository root:

```sh
make deps
export CGO_LDFLAGS="-L$PWD/lib/$(go env GOOS)-$(go env GOARCH)"
picker_library=libonnxruntime.so
if [ "$(go env GOOS)" = darwin ]; then picker_library=libonnxruntime.dylib; fi
export TEST_ONNX_RUNTIME_PATH="$PWD/lib/$(go env GOOS)-$(go env GOARCH)/lib/$picker_library"
go vet ./...
go test ./... -count=1
make build
```

Tests cover HTTP/auth/configuration and one-tier membership. They need native
libraries for linking, but do not download weights. Isolated fallback tests are deferred.

## Real-model smoke

Use [smoke.yaml](../testdata/smoke.yaml): two tiers, no fallback to mask inference errors.
Download the model, then start in one terminal:

```sh
export API_KEY="$(openssl rand -hex 32)"
export CONFIG_PATH="$PWD/testdata/smoke.yaml"
export LISTEN_ADDR=127.0.0.1:8080
make run
```

In another terminal, set the same key and linker settings above, then:

```sh
export CONFIG_PATH="$PWD/testdata/smoke.yaml"
export PICKER_SMOKE_URL=http://127.0.0.1:8080
go test . -run '^TestRealModelHTTP$' -count=1 -v -timeout=10m
```

The check waits up to 180s for health, requires unauthenticated 401 and authenticated
200 JSON, and checks tier/model membership. It does not score selection quality.
For containers, use the same fixture/key with the [quickstart](../README.md#quickstart).

## Selection runner

Requires Python 3.11+ and local loopback access:

```sh
python3 -m unittest discover -s evals -p 'test_*.py' -v
```

These tests start real temporary HTTP subprocesses and replace only the native build step.
They cover scoring, snapshots, response validation, truncated responses, startup failure, startup timeout, SIGTERM cleanup, and surviving child processes.
They do not require native libraries or model weights.

For an actual classifier evaluation with managed startup and shutdown, run:

```sh
python3 evals/run.py
```

The [evaluation guide](../evals/README.md) explains the native prerequisites, isolated configuration, outputs, and exit codes.

## CI

[Workflow](../.github/workflows/ci.yml): native Linux AMD64/ARM64 and macOS ARM64
format/vet/test/build jobs, plus Linux AMD64 service/downloader image builds.
Native jobs also run the Python lifecycle tests without model downloads.
Toolchains, native archives, and action revisions are pinned. No images are published.

Manual dispatch with `real_model` enabled adds native smoke and Compose lifecycle checks.
Each of four runners downloads about 1.6 GiB; ordinary PRs do not download weights.
Temporary keys are masked. Artifacts retain service/smoke logs and model-cache integrity records.
The container check requires graceful exit 0, cached restart without the downloader,
unchanged model hashes/sizes/timestamps, and a valid response using the README default policy.
Fixture requests use two tiers without fallback; default-policy logs must show no inference fallback.

## Verified candidate

[2026-10-28 run](https://github.com/calvintvu/modelpicker/actions/runs/40779512907)
passed all four jobs at `446d1617889d81cc6f639bfdef1daf49e780915c`:

| Target | Environment | Result |
| --- | --- | --- |
| Linux AMD64 native | Ubuntu 24.04.5, GCC 13.3.0 | Build, tests, vet, real-model health/auth/pick passed. |
| Linux ARM64 native | Ubuntu 24.04.5, GCC 13.3.0 | Build, tests, vet, real-model health/auth/pick passed. |
| macOS ARM64 native | macOS 15.7.9, Apple clang 17.0.0 | Build, tests, vet, real-model health/auth/pick passed. |
| Linux AMD64 container | Ubuntu 24.04.5 host, pinned Bookworm image | Fresh download, smoke, shutdown, cache reuse, restart, and default-policy checks passed. |

All use Go 1.27.3, ONNX Runtime 1.29.0, tokenizers v1.27.0, and the README's pinned model revision.
CI's download-only Python is 3.13.17, available on all three native runners.
The downloader image retains its separate Python 3.13.18 pin.
macOS CI targets 15.5 to match the tokenizer archive; only the listed OS versions are verified.
The upstream wrappers still emit a duplicate `-ldl` warning on macOS; checks pass.
These are functional checks, not selection-quality evaluations or load benchmarks.
