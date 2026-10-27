# Third-party licenses and attribution

The router source is licensed under [MIT](LICENSE), copyright 2026 Calvin Vu.
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

Include the following alongside the executable in each binary distribution:

```text
ez-model-router-bin
LICENSE
THIRD_PARTY_NOTICES.md
third_party/
```

For container releases, copy the license and notice bundle into `/usr/share/licenses/ez-model-router/` in the final image.
Retain the runtime's license and notice files whenever distributing the runtime libraries.
For separately distributed model bundles, include the model license text and attribution with the model files.

Container packaging must also retain the base image's package copyright files and applicable source-access information.
The final image's operating-system libraries and package versions must be reviewed when the Dockerfile exists.
System libraries supplied by the user's OS are separate from the native archives reviewed here.

After dependency, model, compiler, or packaging changes, update this inventory and the notices before publishing artifacts.
Preserve upstream license text and copyright notices when updating the bundle.
