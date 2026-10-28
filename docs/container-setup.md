# Container details

Start with the [README quickstart](../README.md#quickstart).
No host Go, Python, or hf CLI is required. Startup does not download weights.

If upgrading a running old checkout, stop its stack before switching names:
`docker compose down --remove-orphans`. This preserves the model directory.

- MODEL_DIR and CONFIG_PATH are host paths; the picker mounts both read-only.
  Only the explicit downloader writes MODEL_DIR. Relative paths use the Compose directory.
- `docker compose down` preserves the model directory and Hugging Face metadata.
  Cached restarts use `docker compose up -d picker`.
- After configuration changes, use `docker compose up -d --force-recreate picker`.
- Linux downloads can be root-owned. For host ownership, create MODEL_DIR first,
  then add `--user "$(id -u):$(id -g)"` to `docker compose run`.
  Files must remain readable by picker UID/GID 65532:65532.
- Shutdown allows 35 seconds for the service's 30-second grace period.
- There is no automatic container healthcheck. `/healthz` reports process health
  after initialization; see the [API contract](api.md#get-healthz).

## Troubleshooting

Check `docker compose logs --tail 50 picker`.

| Failure | Action |
| --- | --- |
| Builder/context error | Select the engine explicitly: `docker --context=NAME compose ...`. |
| Missing API_KEY / 401 | Set a nonempty key; use that same key in the bearer header. Recreate after environment changes. |
| Missing files / failed download | Check paths, permissions, network, and disk; rerun the downloader. |
| Checksum mismatch | Remove only the named damaged file and rerun the downloader. |
| Native-library load error / crash | Match the engine CPU architecture; rebuild pinned dependencies. For native setup, check ONNX_RUNTIME_PATH. |
| Invalid configuration | Fix the field named in logs; check tier counts, descriptions, models, and token fit. |
| Slow initialization / response | Inspect logs; retry health. Increase positive WRITE_TIMEOUT for slow inference, not startup. |

Podman can create empty bind directories despite `create_host_path: false`;
missing required files still prevent serving. Build for the engine's architecture:
emulation can crash or run slowly. Linux containers on macOS use a Linux VM.
Never share rendered Compose configuration containing API_KEY.

Images retain notices under `/usr/share/licenses/modelpicker/`.
For redistribution, follow [third-party notices](../THIRD_PARTY_NOTICES.md) and
[corresponding-source requirements](../third_party/containers/SOURCES.md).
