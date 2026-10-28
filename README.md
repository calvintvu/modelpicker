# ez-model-router

A CPU-only router that classifies a prompt into a configured tier and returns a
random model ID from that tier. It does not call providers or generate completions.

## Quickstart

Requires Git, Docker Engine with Compose v2, curl, OpenSSL, and free port 8080.
Use the engine's CPU architecture; emulation can crash or run slowly.

### 1. Clone the repository

```sh
git clone https://github.com/calvintvu/ez-model-router.git
cd ez-model-router
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

### 4. Start the router

```sh
docker compose up --build -d
```

### 5. Wait for startup

```sh
docker compose logs -f router
```

Model loading can take several minutes. Wait for `server starting on :8080`.
Press Ctrl+C to leave the log view; the router keeps running.

### 6. Check health

```sh
curl -i http://localhost:8080/healthz
```

Expect HTTP 200 with an empty body. Health checks HTTP responsiveness, not inference.

### 7. Route a prompt

Replace `YOUR_API_KEY` with the same key saved in `.env`:

```sh
curl http://localhost:8080/v1/route \
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

Edit `config.yaml`, then recreate the service with `docker compose up -d --force-recreate router`.
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
First complete the clone and API-key steps above. Stop any existing router on port 8080.

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

### 3. Start the router

```sh
./ez-model-router-bin
```

The service reads the key from `.env`. Use the health and routing checks above.
Press Ctrl+C to stop it.

## Platform status

| Target | Status |
| --- | --- |
| macOS ARM64 native | Verified on macOS 27.0.1. |
| Linux AMD64 native/container | Planned; builds pass under emulation, native real-model HTTP remains pending. |
| Linux ARM64 native | Planned; native build/tests pass, native deployment walkthrough remains pending. |
| Linux ARM64 container | Deferred for release; native diagnostic smoke passed. |
| macOS AMD64 | Deferred; no matching pinned runtime archive. |
| Windows AMD64 | Deferred by scope. |

[Startup observations](docs/setup-observations.md) describe one measured machine.
Concurrency safety, overload handling, throughput, and load testing are deferred.
Routing-quality evaluations are also pending. No production capacity is claimed.

## Model and licenses

Uses [receptron/laya-onnx](https://huggingface.co/receptron/laya-onnx/tree/68f27dfe5a27a54fb2b1fefc432f43f972e90868)
at revision `68f27dfe5a27a54fb2b1fefc432f43f972e90868`, including its graph, external weights, and tokenizer.
The same bundle is used across targets.

Source: [MIT](LICENSE). Model: Apache-2.0. Dependency/model attribution and
redistribution requirements: [third-party notices](THIRD_PARTY_NOTICES.md).
Development checks and explicit model smoke: [testing](docs/testing.md).
