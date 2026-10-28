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

## CI

[Workflow](../.github/workflows/ci.yml): native Linux AMD64/ARM64 and macOS ARM64
format/vet/test/build jobs, plus Linux AMD64 service/downloader image builds.
Toolchains, native archives, and action revisions are pinned. No images are published.

Manual dispatch with `real_model` enabled adds native and Compose smoke checks.
Each of four runners downloads about 1.6 GiB; ordinary PRs do not download weights.
Temporary keys are masked. Per-target artifacts retain service and smoke logs.
A defined workflow is not executed evidence: hosted runs and native AMD64 model
verification remain pending. Emulation does not establish native compatibility.
