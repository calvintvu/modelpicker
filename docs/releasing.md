# Release process

The initial release distributes source. Users build locally with Make or Docker Compose.
No prebuilt binaries, registry images, native libraries, or model weights are published.
The pinned dependency and model downloads remain separate setup steps.

Before tagging:

1. Commit public documentation and freeze the application, configuration, model, and dependency revisions.
2. Run CI manually on the candidate branch with `real_model` enabled. Require all four jobs to pass.
3. Check the Linux AMD64 lifecycle results: fresh download, health/authentication/model selection,
   graceful shutdown, cached restart, unchanged model files, and the README default configuration.
4. Record the run URL and tested OS versions in the [testing guide](testing.md).
   Keep unsupported targets and known limitations explicit in the README.
5. Complete the separate selection-quality evaluation and obtain the maintainer's release decision.
6. Tag the approved commit and create a GitHub source release. Notes must identify the model revision,
   verified platforms, `/v1/pick` API, and deferred load validation.

Source archives must include LICENSE, THIRD_PARTY_NOTICES.md, and third_party/.
They must exclude `.env`, downloaded weights, native libraries, and local binaries.

If binary or image releases are added later, include the notice bundle and required
corresponding-source access. Publish artifact checksums and versioned download references.
See [distribution requirements](../THIRD_PARTY_NOTICES.md#distribution-requirements).
