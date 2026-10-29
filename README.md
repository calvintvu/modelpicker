# modelpicker

[![Go version](https://img.shields.io/github/go-mod/go-version/calvintvu/modelpicker/main)](go.mod) [![CI status](https://github.com/calvintvu/modelpicker/actions/workflows/ci.yml/badge.svg)](https://github.com/calvintvu/modelpicker/actions/workflows/ci.yml) [![Source license](https://img.shields.io/github/license/calvintvu/modelpicker)](LICENSE)

A CPU-only model picker that uses a decision model to classify a prompt into a
configured tier, then returns a random model ID from that tier.
It does not proxy LLM calls or generate completions.

## Quickstart

Requires Git, Docker Engine with Compose v2, curl, OpenSSL, and free port 8080.
Use the engine's CPU architecture; emulation can crash or run slowly.

### 1. Clone the repository

```sh
git clone https://github.com/calvintvu/modelpicker.git
cd modelpicker
```

### 2. Set your API key

Copy the settings file:

```sh
cp .env.example .env
chmod 600 .env
```

Generate a key:

```sh
openssl rand -hex 32
```

Paste the generated value into `API_KEY` in `.env`. Keep the key private.
For a slow CPU/VM, also set `WRITE_TIMEOUT=300s`; the default is 60s.

### 3. Download the model

```sh
docker compose run --build --rm model-download
```

Wait for this command to succeed before continuing. The bundle needs about **1.6 GiB**.
Allow **4 GiB** for model/transfer files, plus image/build storage.
No host Go, Python, or Hugging Face CLI is needed. Repeat downloads reuse unchanged files.

### 4. Start the picker

```sh
docker compose up --build -d
```

### 5. Wait for startup

```sh
docker compose logs -f picker
```

Model loading can take several minutes. Wait for `server starting on :8080`.
Press Ctrl+C to leave the log view; the picker keeps running.

### 6. Check health

```sh
curl -i http://localhost:8080/healthz
```

Expect HTTP 200 with an empty body. Health checks HTTP responsiveness, not inference.

### 7. Pick a model for a prompt

Replace `YOUR_API_KEY` with the same key saved in `.env`:

```sh
curl http://localhost:8080/v1/pick \
  -H 'Authorization: Bearer YOUR_API_KEY' \
  -H 'Content-Type: application/json' \
  -d '{"prompt":"Write a Go function."}'
```

Expect JSON such as `{"tier":"code","model":"your-code-model"}`, using IDs from your configuration.
See the [API contract](docs/api.md).

### Stop or restart

Stop and remove the containers:

```sh
docker compose down
```

Start again with cached weights:

```sh
docker compose up -d
```

Normal starts and image rebuilds never download weights. `down` preserves MODEL_DIR
(default `inference/model`); deleting that directory removes the weights.
See [container details](docs/container-setup.md) for custom mounts and troubleshooting.

## Configuration

Edit `config.yaml`, then recreate the service with `docker compose up -d --force-recreate picker`.
A small policy:

```yaml
tiers:
  - name: code
    description: programming
    models: [your-code-model]
  - name: general
    description: everyday questions
    models: [your-general-model]
```

Model IDs are arbitrary strings, not a verified provider catalog.

All fields, environment defaults, and precedence: [configuration reference](docs/configuration.md).

## Native setup

Requires Go 1.27.3, Make, a C/C++ toolchain, curl, tar, shasum, and `hf`.
For the CLI, use `uv tool install huggingface_hub==2.3.0` with uv installed.
First complete the clone and API-key steps above. Stop any existing picker on port 8080.

### 1. Download the model

```sh
make model
```

Reuse an existing MODEL_DIR to avoid another download.

### 2. Build the service

```sh
make build
```

This installs checksum-pinned ONNX Runtime 1.29.0 and tokenizers v1.27.0 for the local target.
Linux uses glibc; Alpine/musl is not verified.

### 3. Start the picker

```sh
./modelpicker-bin
```

The service reads the key from `.env`. Use the health and selection checks above.
Press Ctrl+C to stop it.

## Platform status

| Target | Status |
| --- | --- |
| macOS ARM64 native | Verified on macOS 15.7.9 and 27.0.1; CI build targets macOS 15.5. |
| Linux AMD64 native/container | Verified on Ubuntu 24.04.5; container uses pinned Debian Bookworm. |
| Linux ARM64 native | Verified on Ubuntu 24.04.5. |
| Linux ARM64 container | Deferred for release; native diagnostic smoke passed. |
| macOS AMD64 | Deferred; no matching pinned runtime archive. |
| Windows AMD64 | Deferred by scope. |

Verification includes builds and real-model HTTP checks on native architecture runners:
[CI evidence](https://github.com/calvintvu/modelpicker/actions/runs/40779512907).
Other OS versions are not verified.

[Startup observations](docs/setup-observations.md) describe one measured machine.
Concurrency safety, overload handling, throughput, and load testing are deferred.
Selection-policy evaluations use [labeled prompt suites](evals/README.md). No production capacity is claimed.

## Selection evaluation

With Python 3.11+ and the native build tools installed, run:

```sh
python3 evals/run.py
```

The runner builds and owns a private native server, manages its configuration,
scores the labeled prompts, saves results under `evals/runs/`, and stops the server.
No manual server or API key is required. See the [evaluation guide](evals/README.md)
for prerequisites, cache behavior, outputs, and overrides.

## Model and licenses

Uses [receptron/laya-onnx](https://huggingface.co/receptron/laya-onnx/tree/68f27dfe5a27a54fb2b1fefc432f43f972e90868)
at revision `68f27dfe5a27a54fb2b1fefc432f43f972e90868`, including its graph, external weights, and tokenizer.
The same bundle is used across targets.

Source: [MIT](LICENSE). Model: Apache-2.0. Dependency/model attribution and
redistribution requirements: [third-party notices](THIRD_PARTY_NOTICES.md).
Development checks and explicit model smoke: [testing](docs/testing.md).

The initial release format is source with local Make or Docker Compose builds.
See the [release process](docs/releasing.md); prebuilt binaries and registry images are outside this release.
