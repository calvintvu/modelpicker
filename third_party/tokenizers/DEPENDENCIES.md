# Tokenizer dependency license inventory

Pinned source: https://github.com/daulet/tokenizers/tree/v1.27.0

The Go module/native archive release is v1.27.0. Its FFI crate declares version 1.26.0.
This inventory covers the union of normal and build dependency closures for aarch64-apple-darwin, x86_64-unknown-linux-gnu, and aarch64-unknown-linux-gnu.
Development-only edges and the separate WebAssembly workspace root are excluded. Build dependencies are conservatively retained.
The original upstream Cargo.lock is copied alongside this inventory. No project dependency versions were changed.

License expressions below reproduce package metadata. LICENSES.txt preserves supplied license, copyright, and notice files verbatim under component headers, including alternate license texts and nested Oniguruma notices.
The repository-level MIT license covers tokenizers-ffi. tiktoken-rs includes its workspace license and the vendored OpenAI tiktoken license. number_prefix excludes its LICENCE from its crate package; its license was retrieved at source commit eb6ebd215d50df1b199737f0356b988bebaedc84.
Rust standard-library notices are stored separately under ../rust/.

| Package | Version | Declared license | Preserved files |
| --- | --- | --- | --- |
| `ahash` | `0.8.12` | MIT OR Apache-2.0 | 2 |
| `aho-corasick` | `1.1.3` | Unlicense OR MIT | 2 |
| `anyhow` | `1.0.98` | MIT OR Apache-2.0 | 2 |
| `base64` | `0.13.1` | MIT/Apache-2.0 | 2 |
| `base64` | `0.22.1` | MIT OR Apache-2.0 | 2 |
| `bit-set` | `0.8.0` | Apache-2.0 OR MIT | 2 |
| `bit-vec` | `0.8.0` | Apache-2.0 OR MIT | 2 |
| `bitflags` | `2.9.1` | MIT OR Apache-2.0 | 2 |
| `bstr` | `1.12.0` | MIT OR Apache-2.0 | 3 |
| `castaway` | `0.2.4` | MIT | 1 |
| `cc` | `1.2.26` | MIT OR Apache-2.0 | 2 |
| `cfg-if` | `1.0.1` | MIT OR Apache-2.0 | 2 |
| `compact_str` | `0.9.0` | MIT | 1 |
| `console` | `0.15.11` | MIT | 1 |
| `crossbeam-deque` | `0.8.6` | MIT OR Apache-2.0 | 2 |
| `crossbeam-epoch` | `0.9.18` | MIT OR Apache-2.0 | 2 |
| `crossbeam-utils` | `0.8.21` | MIT OR Apache-2.0 | 2 |
| `darling` | `0.20.11` | MIT | 1 |
| `darling_core` | `0.20.11` | MIT | 1 |
| `darling_macro` | `0.20.11` | MIT | 1 |
| `dary_heap` | `0.3.7` | MIT OR Apache-2.0 | 2 |
| `derive_builder` | `0.20.2` | MIT OR Apache-2.0 | 2 |
| `derive_builder_core` | `0.20.2` | MIT OR Apache-2.0 | 2 |
| `derive_builder_macro` | `0.20.2` | MIT OR Apache-2.0 | 2 |
| `either` | `1.15.0` | MIT OR Apache-2.0 | 2 |
| `esaxx-rs` | `0.1.10` | Apache-2.0 | 1 |
| `fancy-regex` | `0.14.0` | MIT | 1 |
| `fancy-regex` | `0.16.1` | MIT | 1 |
| `fnv` | `1.0.7` | Apache-2.0 / MIT | 2 |
| `getrandom` | `0.3.3` | MIT OR Apache-2.0 | 2 |
| `ident_case` | `1.0.1` | MIT/Apache-2.0 | 1 |
| `indicatif` | `0.17.11` | MIT | 1 |
| `itertools` | `0.14.0` | MIT OR Apache-2.0 | 2 |
| `itoa` | `1.0.15` | MIT OR Apache-2.0 | 2 |
| `lazy_static` | `1.5.0` | MIT OR Apache-2.0 | 2 |
| `libc` | `0.2.172` | MIT OR Apache-2.0 | 2 |
| `log` | `0.4.27` | MIT OR Apache-2.0 | 2 |
| `macro_rules_attribute` | `0.2.2` | Apache-2.0 OR MIT OR Zlib | 3 |
| `macro_rules_attribute-proc_macro` | `0.2.2` | Apache-2.0 OR MIT OR Zlib | 3 |
| `memchr` | `2.7.4` | Unlicense OR MIT | 2 |
| `minimal-lexical` | `0.2.1` | MIT/Apache-2.0 | 3 |
| `monostate` | `0.1.14` | MIT OR Apache-2.0 | 2 |
| `monostate-impl` | `0.1.14` | MIT OR Apache-2.0 | 2 |
| `nom` | `7.1.3` | MIT | 1 |
| `number_prefix` | `0.4.0` | MIT | 1 |
| `once_cell` | `1.21.3` | MIT OR Apache-2.0 | 2 |
| `onig` | `6.5.1` | MIT | 1 |
| `onig_sys` | `69.9.1` | MIT | 2 |
| `paste` | `1.0.15` | MIT OR Apache-2.0 | 2 |
| `pkg-config` | `0.3.32` | MIT OR Apache-2.0 | 2 |
| `portable-atomic` | `1.11.1` | Apache-2.0 OR MIT | 2 |
| `ppv-lite86` | `0.2.21` | MIT OR Apache-2.0 | 2 |
| `proc-macro2` | `1.0.95` | MIT OR Apache-2.0 | 2 |
| `quote` | `1.0.40` | MIT OR Apache-2.0 | 2 |
| `rand` | `0.9.2` | MIT OR Apache-2.0 | 3 |
| `rand_chacha` | `0.9.0` | MIT OR Apache-2.0 | 3 |
| `rand_core` | `0.9.3` | MIT OR Apache-2.0 | 3 |
| `rayon` | `1.10.0` | MIT OR Apache-2.0 | 2 |
| `rayon-cond` | `0.4.0` | Apache-2.0/MIT | 2 |
| `rayon-core` | `1.12.1` | MIT OR Apache-2.0 | 2 |
| `regex` | `1.11.1` | MIT OR Apache-2.0 | 2 |
| `regex-automata` | `0.4.9` | MIT OR Apache-2.0 | 2 |
| `regex-syntax` | `0.8.5` | MIT OR Apache-2.0 | 3 |
| `rustc-hash` | `1.1.0` | Apache-2.0/MIT | 2 |
| `rustversion` | `1.0.21` | MIT OR Apache-2.0 | 2 |
| `ryu` | `1.0.20` | Apache-2.0 OR BSL-1.0 | 2 |
| `serde` | `1.0.219` | MIT OR Apache-2.0 | 2 |
| `serde_derive` | `1.0.219` | MIT OR Apache-2.0 | 2 |
| `serde_json` | `1.0.140` | MIT OR Apache-2.0 | 2 |
| `shlex` | `1.3.0` | MIT OR Apache-2.0 | 2 |
| `smallvec` | `1.15.1` | MIT OR Apache-2.0 | 2 |
| `spm_precompiled` | `0.1.4` | Apache-2.0 | 1 |
| `static_assertions` | `1.1.0` | MIT OR Apache-2.0 | 2 |
| `strsim` | `0.11.1` | MIT | 1 |
| `syn` | `2.0.102` | MIT OR Apache-2.0 | 2 |
| `thiserror` | `2.0.16` | MIT OR Apache-2.0 | 2 |
| `thiserror-impl` | `2.0.16` | MIT OR Apache-2.0 | 2 |
| `tiktoken-rs` | `0.8.0` | MIT | 2 |
| `tokenizers` | `0.22.0` | Apache-2.0 | 1 |
| `tokenizers-ffi` | `1.26.0` | MIT (repository license) | 1 |
| `unicode-ident` | `1.0.18` | (MIT OR Apache-2.0) AND Unicode-3.0 | 3 |
| `unicode-normalization-alignments` | `0.1.12` | MIT/Apache-2.0 | 3 |
| `unicode-segmentation` | `1.12.0` | MIT OR Apache-2.0 | 3 |
| `unicode-width` | `0.2.1` | MIT OR Apache-2.0 | 3 |
| `unicode_categories` | `0.1.1` | MIT OR Apache-2.0 | 2 |
| `version_check` | `0.9.5` | MIT/Apache-2.0 | 2 |
| `zerocopy` | `0.8.25` | BSD-2-Clause OR Apache-2.0 OR MIT | 3 |
