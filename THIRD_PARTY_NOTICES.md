# Third-party licenses and attribution

The picker source is licensed under [MIT](LICENSE), copyright 2026 Calvin Vu.
Third-party code, model files, and native libraries retain their own licenses and notices.

## Model bundle

The default model is [receptron/laya-onnx](https://huggingface.co/receptron/laya-onnx/tree/68f27dfe5a27a54fb2b1fefc432f43f972e90868), revision `68f27dfe5a27a54fb2b1fefc432f43f972e90868`.
Its model card declares Apache-2.0 for the exported model bundle and attributes the weights to Convai Innovations.
It includes the ONNX graph, external weights, configuration, and checkpoint tokenizer files.

Model attribution:

- Laya weights: Convai Innovations, [convaiinnovations/laya](https://huggingface.co/convaiinnovations/laya).
- ONNX export: Receptron, [receptron/laya-onnx](https://huggingface.co/receptron/laya-onnx).
- ModernBERT: Answer.AI, LightOn, and contributors, [answerdotai/ModernBERT-large](https://huggingface.co/answerdotai/ModernBERT-large).

The Apache-2.0 license text is included in [third_party/Apache-2.0.txt](third_party/Apache-2.0.txt).
The pinned export repository supplies no separate LICENSE or NOTICE file. Its model card supplies the license declaration and attribution.
The exporter project's separate MIT license applies to its code. The model weights retain Apache-2.0 terms.

For model redistribution, include the Apache-2.0 text and these model attributions with the bundle.
Retain any additional upstream notices and identify modifications to redistributed model files.

## Code and native libraries

| Component | Version | Terms and preserved notices |
| --- | --- | --- |
| `github.com/daulet/tokenizers` Go bindings | v1.27.0 | MIT; copyright 2023 Daulet Zhanguzin. [Original license](third_party/go/LICENSES.txt). |
| `github.com/yalue/onnxruntime_go` | v1.36.0 | MIT; copyright 2023 Nathan Otterness. [Original license](third_party/go/LICENSES.txt). |
| `github.com/joho/godotenv` | v1.5.1 | MIT; copyright 2013 John Barton. [Original license](third_party/go/LICENSES.txt). |
| `go.yaml.in/yaml/v3` | v3.0.5 | MIT and Apache-2.0 apply to different files. Preserve its original LICENSE and NOTICE, both in [Go notices](third_party/go/LICENSES.txt). |
| Go runtime and standard-library material | 1.27.3 build toolchain | BSD-3-Clause and the supplied patent grant. [Go and vendored notices](third_party/go/LICENSES.txt). |
| ONNX Runtime CPU libraries | 1.29.0 | [Microsoft MIT license](third_party/onnxruntime/LICENSE) and the complete supplied [ThirdPartyNotices.txt](third_party/onnxruntime/ThirdPartyNotices.txt). |
| Tokenizer native static library | daulet/tokenizers v1.27.0 | Repository MIT license plus the Rust dependency terms listed in the [inventory](third_party/tokenizers/DEPENDENCIES.md). [Original license and notice texts](third_party/tokenizers/LICENSES.txt). |
| Rust standard-library material embedded in the tokenizer | Rust 1.94.1, commit `e408947bfd200af42db322daf0fadfe7e26d3bd1` | [Copyright information](third_party/rust/COPYRIGHT), [MIT](third_party/rust/LICENSE-MIT), [Apache-2.0](third_party/rust/LICENSE-APACHE), and original [standard-library notices](third_party/rust/COPYRIGHT-library.html). |

The tokenizer inventory covers 87 packages, including the FFI crate and its normal/build dependencies across the three planned native targets.
It conservatively includes build dependencies and supplied alternate license texts. Third-party license files are preserved verbatim under component headers.
The Rust notices also cover the embedded standard-library dependencies. The runtime's supplied notice file is preserved in full, including optional-component notices.

## Distribution requirements

### Optional model downloader image

`docker/model-download.Dockerfile` uses Python 3.13.18 and Hugging Face CLI
(`huggingface_hub`) 2.3.0. Its Python dependencies are pinned with artifact hashes in
`docker/model-download-requirements.txt`. They are separate from the Go service image.

The installed Linux dependency terms are:

| Components | Terms |
| --- | --- |
| `huggingface_hub` 2.3.0, `hf-xet` 1.7.0 | Apache-2.0 |
| `anyio` 4.15.1, `filelock` 4.2.4, `h11` 0.16.0, PyYAML 6.0.3, `truststore` 0.10.5 | MIT |
| `click` 8.5.0, `fsspec` 2026.9.0, `httpcore2` 2.16.0, `httpx2` 2.16.0, `idna` 3.20 | BSD-3-Clause |
| `packaging` 26.3 | Apache-2.0 OR BSD-2-Clause |
| `tqdm` 4.70.1 | MPL-2.0 AND MIT |
| `typing_extensions` 4.16.0 | PSF-2.0 |

The image preserves each wheel's original license files under
`/usr/local/lib/python3.13/site-packages/*.dist-info/licenses/`.
It also preserves the base image's Python license at
`/usr/local/lib/python3.13/LICENSE.txt`, pip 26.2.2's MIT and vendored notices,
and Debian package copyright files under `/usr/share/doc/`.
Retain these files when distributing or modifying the downloader image.
For `tqdm`, provide access to the corresponding unmodified source when distributing
the image; its versioned source distribution is available from
[PyPI](https://pypi.org/project/tqdm/4.70.1/#files).
The native `hf-xet` normal/build dependency inventory, original notices, upstream
lockfile, and matching Rust 1.99.0 notices are included under
[third_party/hf-xet](third_party/hf-xet/DEPENDENCIES.md).
[Container source access](third_party/containers/SOURCES.md) records the complete
Debian package inventories and corresponding source archive locations.
Those reviewed inventories use the former image label `router` for the picker service.
When publishing images, provide recipients corresponding source access under the
applicable terms; these inventories do not create a written source offer.

### Picker binary and image

Include the following alongside the executable in each binary distribution:

```text
modelpicker-bin
LICENSE
THIRD_PARTY_NOTICES.md
third_party/
```

For container releases, copy the license and notice bundle into `/usr/share/licenses/modelpicker/` in the final image.
Retain the runtime's license and notice files whenever distributing the runtime libraries.
For separately distributed model bundles, include the model license text and attribution with the model files.

Container packaging retains the base image's package copyright files and applicable source-access information.
The pinned runtime and downloader base package versions are recorded in the
[container inventories](third_party/containers/SOURCES.md), including exact source archive references.
System libraries supplied by the user's OS are separate from the native archives reviewed here.

After dependency, model, compiler, or packaging changes, update this inventory and the notices before publishing artifacts.
Preserve upstream license text and copyright notices when updating the bundle.
