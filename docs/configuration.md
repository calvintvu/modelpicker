# Configuration

Edit `config.yaml`; use the [small example](../README.md#configuration).
Recreate containers with `docker compose up -d --force-recreate picker`.
Restart a native process after changes; there is no hot reload.

| YAML field | Rule / default |
| --- | --- |
| `classifier.question` | Default: `What kind of request is this?` |
| `classifier.temperature` | `0` uses tier-count calibration; a positive finite value overrides it. This is not model sampling. |
| `default_tier` | Optional tier name used only for inference errors, not low confidence. Startup errors still stop the process. |
| `tiers` | List of 1-10 tiers. |
| `tiers[].name` | Unique, nonblank label returned as `tier`. |
| `tiers[].description` | Nonblank for multiple tiers; optional for one tier. |
| `tiers[].models` | Nonempty list of nonblank IDs; one entry is chosen at random. |

Unknown YAML fields fail validation. Multi-tier question/options must fit the
512-token classifier budget with space for a prompt token; long prompts are truncated.

| Environment | Native default | Compose behavior |
| --- | --- | --- |
| `API_KEY` | Required | Required |
| `CONFIG_PATH` | `./config.yaml` | Host file mounted at `/app/config.yaml` |
| `MODEL_DIR` | `./inference/model` | Host directory mounted at `/models` |
| `LISTEN_ADDR` | `:8080` | Fixed `:8080`, published on port 8080 |
| `WRITE_TIMEOUT` | `60s`; positive Go duration | Same |
| `ONNX_RUNTIME_PATH` | `lib/<GOOS>-<GOARCH>/lib/libonnxruntime.so` or `.dylib` | Packaged library at `/usr/local/lib/libonnxruntime.so` |

Existing environment values override optional `.env` values, then defaults apply.
Compose uses `.env` for interpolation and fixes container paths as shown above.
Make does not read `.env`; export MODEL_DIR before `make model` for a custom native directory.
Native paths use the working directory; Compose paths use the Compose directory.
The API key is separate from model/provider IDs.
