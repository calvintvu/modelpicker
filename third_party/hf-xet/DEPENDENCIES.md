# hf-xet native dependency notices

Version: 1.7.0. Reviewed 2026-10-27.

This inventory covers 231 normal/build packages reachable from hf_xet for
x86_64-unknown-linux-gnu and aarch64-unknown-linux-gnu, with the published
pyo3/extension-module feature. Development-only dependencies are excluded.
Build dependencies and alternate license texts are conservatively preserved.
The source distribution SHA-256 is
`d406ec79053c0871817f700c2ac8c36ba0d87f9c34b7458b0f0063bb218b0466`.
The [upstream lockfile](Cargo.lock) is preserved without edits.

[Original notices](LICENSES.txt) preserve 420 license/copyright/notice files,
including nested vendored AWS-LC and ring notices.
Local xet packages inherit the source distribution's Apache-2.0 license.
Registry package sources can be downloaded from the links below and verified
against package checksums in the upstream lockfile. Local package sources are
in the [pinned source distribution](https://files.pythonhosted.org/packages/9e/27/06d899ea7bd721d272f84aac98bdb238de98af4cc767a69056d967d68c71/hf_xet-1.7.0.tar.gz).
Preserve these notices and supply corresponding source access for MPL-2.0 material
when distributing the image; do not remove this source inventory.

Both reviewed Linux wheels embed Rust compiler commit
`b940084d7eb6a299eb4bfeb8e34901bc051e7ac4` (Rust 1.99.0).
The matching [Rust notices](rust/COPYRIGHT-library.html), COPYRIGHT, MIT and
Apache-2.0 texts are preserved from the official compiler archive.
Its SHA-256 is `334a66714ca316d71bbe5efe44f71762d0b276cea73be2987374693a837f7450`.
Rust source: https://github.com/rust-lang/rust/tree/b940084d7eb6a299eb4bfeb8e34901bc051e7ac4

| Package | Version | Declared terms | Source |
| --- | --- | --- | --- |
| aho-corasick | 1.1.4 | Unlicense OR MIT | [source](https://crates.io/api/v1/crates/aho-corasick/1.1.4/download) |
| anyhow | 1.0.103 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/anyhow/1.0.103/download) |
| approx | 0.5.1 | Apache-2.0 | [source](https://crates.io/api/v1/crates/approx/0.5.1/download) |
| arrayref | 0.3.9 | BSD-2-Clause | [source](https://crates.io/api/v1/crates/arrayref/0.3.9/download) |
| arrayvec | 0.7.6 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/arrayvec/0.7.6/download) |
| async-trait | 0.1.89 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/async-trait/0.1.89/download) |
| atomic-waker | 1.1.2 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/atomic-waker/1.1.2/download) |
| autocfg | 1.5.0 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/autocfg/1.5.0/download) |
| aws-lc-rs | 1.18.1 | ISC AND (Apache-2.0 OR ISC) | [source](https://crates.io/api/v1/crates/aws-lc-rs/1.18.1/download) |
| aws-lc-sys | 0.45.0 | ISC AND (Apache-2.0 OR ISC) AND Apache-2.0 AND MIT AND BSD-3-Clause AND (Apache-2.0 OR ISC OR MIT) AND (Apache-2.0 OR ISC OR MIT-0) | [source](https://crates.io/api/v1/crates/aws-lc-sys/0.45.0/download) |
| base64 | 0.22.1 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/base64/0.22.1/download) |
| bitflags | 2.11.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/bitflags/2.11.0/download) |
| blake3 | 1.8.3 | CC0-1.0 OR Apache-2.0 OR Apache-2.0 WITH LLVM-exception | [source](https://crates.io/api/v1/crates/blake3/1.8.3/download) |
| block-buffer | 0.12.1 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/block-buffer/0.12.1/download) |
| bstr | 1.12.1 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/bstr/1.12.1/download) |
| bytemuck | 1.25.0 | Zlib OR Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/bytemuck/1.25.0/download) |
| bytes | 1.11.1 | MIT | [source](https://crates.io/api/v1/crates/bytes/1.11.1/download) |
| cc | 1.2.57 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/cc/1.2.57/download) |
| cfg-if | 0.1.10 | MIT/Apache-2.0 | [source](https://crates.io/api/v1/crates/cfg-if/0.1.10/download) |
| cfg-if | 1.0.4 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/cfg-if/1.0.4/download) |
| cfg_aliases | 0.2.1 | MIT | [source](https://crates.io/api/v1/crates/cfg_aliases/0.2.1/download) |
| chacha20 | 0.10.2 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/chacha20/0.10.2/download) |
| chrono | 0.4.44 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/chrono/0.4.44/download) |
| cmake | 0.1.57 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/cmake/0.1.57/download) |
| colored | 3.1.1 | MPL-2.0 | [source](https://crates.io/api/v1/crates/colored/3.1.1/download) |
| const-oid | 0.10.2 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/const-oid/0.10.2/download) |
| const-str | 1.1.0 | MIT | [source](https://crates.io/api/v1/crates/const-str/1.1.0/download) |
| const_panic | 0.2.15 | Zlib | [source](https://crates.io/api/v1/crates/const_panic/0.2.15/download) |
| constant_time_eq | 0.4.2 | CC0-1.0 OR MIT-0 OR Apache-2.0 | [source](https://crates.io/api/v1/crates/constant_time_eq/0.4.2/download) |
| countio | 0.3.0 | MIT | [source](https://crates.io/api/v1/crates/countio/0.3.0/download) |
| cpufeatures | 0.2.17 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/cpufeatures/0.2.17/download) |
| cpufeatures | 0.3.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/cpufeatures/0.3.0/download) |
| crc32fast | 1.5.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/crc32fast/1.5.0/download) |
| crossbeam-channel | 0.5.15 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/crossbeam-channel/0.5.15/download) |
| crossbeam-utils | 0.8.21 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/crossbeam-utils/0.8.21/download) |
| crypto-common | 0.2.2 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/crypto-common/0.2.2/download) |
| ctor | 1.0.3 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/ctor/1.0.3/download) |
| deranged | 0.5.8 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/deranged/0.5.8/download) |
| digest | 0.11.3 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/digest/0.11.3/download) |
| dirs | 6.0.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/dirs/6.0.0/download) |
| dirs-sys | 0.5.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/dirs-sys/0.5.0/download) |
| displaydoc | 0.2.5 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/displaydoc/0.2.5/download) |
| dunce | 1.0.5 | CC0-1.0 OR MIT-0 OR Apache-2.0 | [source](https://crates.io/api/v1/crates/dunce/1.0.5/download) |
| either | 1.15.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/either/1.15.0/download) |
| errno | 0.3.14 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/errno/0.3.14/download) |
| fastrand | 2.3.0 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/fastrand/2.3.0/download) |
| find-msvc-tools | 0.1.9 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/find-msvc-tools/0.1.9/download) |
| form_urlencoded | 1.2.2 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/form_urlencoded/1.2.2/download) |
| fs_extra | 1.3.0 | MIT | [source](https://crates.io/api/v1/crates/fs_extra/1.3.0/download) |
| futures | 0.3.32 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/futures/0.3.32/download) |
| futures-channel | 0.3.32 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/futures-channel/0.3.32/download) |
| futures-core | 0.3.32 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/futures-core/0.3.32/download) |
| futures-executor | 0.3.32 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/futures-executor/0.3.32/download) |
| futures-io | 0.3.32 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/futures-io/0.3.32/download) |
| futures-macro | 0.3.32 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/futures-macro/0.3.32/download) |
| futures-sink | 0.3.32 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/futures-sink/0.3.32/download) |
| futures-task | 0.3.32 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/futures-task/0.3.32/download) |
| futures-util | 0.3.32 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/futures-util/0.3.32/download) |
| gearhash | 0.1.3 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/gearhash/0.1.3/download) |
| getrandom | 0.2.17 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/getrandom/0.2.17/download) |
| getrandom | 0.3.4 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/getrandom/0.3.4/download) |
| getrandom | 0.4.2 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/getrandom/0.4.2/download) |
| git-version | 0.3.9 | BSD-2-Clause | [source](https://crates.io/api/v1/crates/git-version/0.3.9/download) |
| git-version-macro | 0.3.9 | BSD-2-Clause | [source](https://crates.io/api/v1/crates/git-version-macro/0.3.9/download) |
| heapify | 0.2.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/heapify/0.2.0/download) |
| heck | 0.5.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/heck/0.5.0/download) |
| hf-xet | 1.7.0 | Apache-2.0 | [source](https://pypi.org/project/hf-xet/1.7.0/#files) |
| hf_xet | 1.7.0 | Apache-2.0 | [source](https://pypi.org/project/hf-xet/1.7.0/#files) |
| http | 1.4.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/http/1.4.0/download) |
| http-body | 1.0.1 | MIT | [source](https://crates.io/api/v1/crates/http-body/1.0.1/download) |
| http-body-util | 0.1.3 | MIT | [source](https://crates.io/api/v1/crates/http-body-util/0.1.3/download) |
| httparse | 1.10.1 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/httparse/1.10.1/download) |
| humantime | 2.3.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/humantime/2.3.0/download) |
| hybrid-array | 0.4.13 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/hybrid-array/0.4.13/download) |
| hyper | 1.8.1 | MIT | [source](https://crates.io/api/v1/crates/hyper/1.8.1/download) |
| hyper-rustls | 0.27.7 | Apache-2.0 OR ISC OR MIT | [source](https://crates.io/api/v1/crates/hyper-rustls/0.27.7/download) |
| hyper-util | 0.1.20 | MIT | [source](https://crates.io/api/v1/crates/hyper-util/0.1.20/download) |
| iana-time-zone | 0.1.65 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/iana-time-zone/0.1.65/download) |
| icu_collections | 2.1.1 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/icu_collections/2.1.1/download) |
| icu_locale_core | 2.1.1 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/icu_locale_core/2.1.1/download) |
| icu_normalizer | 2.1.1 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/icu_normalizer/2.1.1/download) |
| icu_normalizer_data | 2.1.1 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/icu_normalizer_data/2.1.1/download) |
| icu_properties | 2.1.2 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/icu_properties/2.1.2/download) |
| icu_properties_data | 2.1.2 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/icu_properties_data/2.1.2/download) |
| icu_provider | 2.1.1 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/icu_provider/2.1.1/download) |
| idna | 1.1.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/idna/1.1.0/download) |
| idna_adapter | 1.2.1 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/idna_adapter/1.2.1/download) |
| inventory | 0.3.22 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/inventory/0.3.22/download) |
| ipnet | 2.12.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/ipnet/2.12.0/download) |
| iri-string | 0.7.11 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/iri-string/0.7.11/download) |
| itertools | 0.14.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/itertools/0.14.0/download) |
| itoa | 1.0.18 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/itoa/1.0.18/download) |
| jobserver | 0.1.34 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/jobserver/0.1.34/download) |
| konst | 0.4.3 | Zlib | [source](https://crates.io/api/v1/crates/konst/0.4.3/download) |
| konst_proc_macros | 0.4.1 | Zlib | [source](https://crates.io/api/v1/crates/konst_proc_macros/0.4.1/download) |
| lazy_static | 1.5.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/lazy_static/1.5.0/download) |
| libc | 0.2.183 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/libc/0.2.183/download) |
| link-section | 0.15.0 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/link-section/0.15.0/download) |
| linktime-proc-macro | 0.1.0 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/linktime-proc-macro/0.1.0/download) |
| linux-raw-sys | 0.12.1 | Apache-2.0 WITH LLVM-exception OR Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/linux-raw-sys/0.12.1/download) |
| litemap | 0.8.1 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/litemap/0.8.1/download) |
| log | 0.4.29 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/log/0.4.29/download) |
| lru-slab | 0.1.2 | MIT OR Apache-2.0 OR Zlib | [source](https://crates.io/api/v1/crates/lru-slab/0.1.2/download) |
| lz4_flex | 0.13.0 | MIT | [source](https://crates.io/api/v1/crates/lz4_flex/0.13.0/download) |
| matchers | 0.2.0 | MIT | [source](https://crates.io/api/v1/crates/matchers/0.2.0/download) |
| memchr | 2.8.0 | Unlicense OR MIT | [source](https://crates.io/api/v1/crates/memchr/2.8.0/download) |
| mio | 1.1.1 | MIT | [source](https://crates.io/api/v1/crates/mio/1.1.1/download) |
| more-asserts | 0.3.1 | Unlicense OR MIT OR Apache-2.0 OR CC0-1.0 | [source](https://crates.io/api/v1/crates/more-asserts/0.3.1/download) |
| nu-ansi-term | 0.50.3 | MIT | [source](https://crates.io/api/v1/crates/nu-ansi-term/0.50.3/download) |
| num-conv | 0.2.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/num-conv/0.2.0/download) |
| num-traits | 0.2.19 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/num-traits/0.2.19/download) |
| once_cell | 1.21.4 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/once_cell/1.21.4/download) |
| oneshot | 0.1.13 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/oneshot/0.1.13/download) |
| openssl-probe | 0.2.1 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/openssl-probe/0.2.1/download) |
| option-ext | 0.2.0 | MPL-2.0 | [source](https://crates.io/api/v1/crates/option-ext/0.2.0/download) |
| os_str_bytes | 6.6.1 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/os_str_bytes/6.6.1/download) |
| percent-encoding | 2.3.2 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/percent-encoding/2.3.2/download) |
| pin-project | 1.1.11 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/pin-project/1.1.11/download) |
| pin-project-internal | 1.1.11 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/pin-project-internal/1.1.11/download) |
| pin-project-lite | 0.2.17 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/pin-project-lite/0.2.17/download) |
| pin-utils | 0.1.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/pin-utils/0.1.0/download) |
| pkg-config | 0.3.32 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/pkg-config/0.3.32/download) |
| potential_utf | 0.1.4 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/potential_utf/0.1.4/download) |
| powerfmt | 0.2.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/powerfmt/0.2.0/download) |
| ppv-lite86 | 0.2.21 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/ppv-lite86/0.2.21/download) |
| proc-macro2 | 1.0.106 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/proc-macro2/1.0.106/download) |
| pyo3 | 0.29.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/pyo3/0.29.0/download) |
| pyo3-build-config | 0.29.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/pyo3-build-config/0.29.0/download) |
| pyo3-ffi | 0.29.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/pyo3-ffi/0.29.0/download) |
| pyo3-macros | 0.29.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/pyo3-macros/0.29.0/download) |
| pyo3-macros-backend | 0.29.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/pyo3-macros-backend/0.29.0/download) |
| quinn | 0.11.9 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/quinn/0.11.9/download) |
| quinn-proto | 0.11.15 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/quinn-proto/0.11.15/download) |
| quinn-udp | 0.5.14 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/quinn-udp/0.5.14/download) |
| quote | 1.0.45 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/quote/1.0.45/download) |
| rand | 0.10.1 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/rand/0.10.1/download) |
| rand | 0.9.4 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/rand/0.9.4/download) |
| rand_chacha | 0.9.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/rand_chacha/0.9.0/download) |
| rand_core | 0.10.1 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/rand_core/0.10.1/download) |
| rand_core | 0.9.5 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/rand_core/0.9.5/download) |
| redb | 3.1.1 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/redb/3.1.1/download) |
| regex | 1.12.3 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/regex/1.12.3/download) |
| regex-automata | 0.4.14 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/regex-automata/0.4.14/download) |
| regex-syntax | 0.8.10 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/regex-syntax/0.8.10/download) |
| reqwest | 0.13.2 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/reqwest/0.13.2/download) |
| reqwest-middleware | 0.5.1 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/reqwest-middleware/0.5.1/download) |
| ring | 0.17.14 | Apache-2.0 AND ISC | [source](https://crates.io/api/v1/crates/ring/0.17.14/download) |
| rustc-hash | 2.1.1 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/rustc-hash/2.1.1/download) |
| rustix | 1.1.4 | Apache-2.0 WITH LLVM-exception OR Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/rustix/1.1.4/download) |
| rustls | 0.23.45 | Apache-2.0 OR ISC OR MIT | [source](https://crates.io/api/v1/crates/rustls/0.23.45/download) |
| rustls-native-certs | 0.8.3 | Apache-2.0 OR ISC OR MIT | [source](https://crates.io/api/v1/crates/rustls-native-certs/0.8.3/download) |
| rustls-pki-types | 1.14.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/rustls-pki-types/1.14.0/download) |
| rustls-platform-verifier | 0.6.2 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/rustls-platform-verifier/0.6.2/download) |
| rustls-webpki | 0.103.15 | ISC | [source](https://crates.io/api/v1/crates/rustls-webpki/0.103.15/download) |
| safe-transmute | 0.11.3 | MIT | [source](https://crates.io/api/v1/crates/safe-transmute/0.11.3/download) |
| serde | 1.0.228 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/serde/1.0.228/download) |
| serde_core | 1.0.228 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/serde_core/1.0.228/download) |
| serde_derive | 1.0.228 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/serde_derive/1.0.228/download) |
| serde_json | 1.0.149 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/serde_json/1.0.149/download) |
| serde_repr | 0.1.20 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/serde_repr/0.1.20/download) |
| sha2 | 0.11.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/sha2/0.11.0/download) |
| sharded-slab | 0.1.7 | MIT | [source](https://crates.io/api/v1/crates/sharded-slab/0.1.7/download) |
| shellexpand | 3.1.2 | MIT/Apache-2.0 | [source](https://crates.io/api/v1/crates/shellexpand/3.1.2/download) |
| shlex | 1.3.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/shlex/1.3.0/download) |
| signal-hook | 0.3.18 | Apache-2.0/MIT | [source](https://crates.io/api/v1/crates/signal-hook/0.3.18/download) |
| signal-hook-registry | 1.4.8 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/signal-hook-registry/1.4.8/download) |
| slab | 0.4.12 | MIT | [source](https://crates.io/api/v1/crates/slab/0.4.12/download) |
| smallvec | 1.15.1 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/smallvec/1.15.1/download) |
| socket2 | 0.6.3 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/socket2/0.6.3/download) |
| stable_deref_trait | 1.2.1 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/stable_deref_trait/1.2.1/download) |
| static_assertions | 1.1.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/static_assertions/1.1.0/download) |
| statrs | 0.18.0 | MIT | [source](https://crates.io/api/v1/crates/statrs/0.18.0/download) |
| subtle | 2.6.1 | BSD-3-Clause | [source](https://crates.io/api/v1/crates/subtle/2.6.1/download) |
| syn | 2.0.117 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/syn/2.0.117/download) |
| sync_wrapper | 1.0.2 | Apache-2.0 | [source](https://crates.io/api/v1/crates/sync_wrapper/1.0.2/download) |
| synstructure | 0.13.2 | MIT | [source](https://crates.io/api/v1/crates/synstructure/0.13.2/download) |
| sysinfo | 0.39.6 | MIT | [source](https://crates.io/api/v1/crates/sysinfo/0.39.6/download) |
| target-lexicon | 0.13.5 | Apache-2.0 WITH LLVM-exception | [source](https://crates.io/api/v1/crates/target-lexicon/0.13.5/download) |
| tempfile | 3.27.0 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/tempfile/3.27.0/download) |
| thiserror | 2.0.18 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/thiserror/2.0.18/download) |
| thiserror-impl | 2.0.18 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/thiserror-impl/2.0.18/download) |
| thread_local | 1.1.9 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/thread_local/1.1.9/download) |
| time | 0.3.47 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/time/0.3.47/download) |
| time-core | 0.1.8 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/time-core/0.1.8/download) |
| time-macros | 0.2.27 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/time-macros/0.2.27/download) |
| tinystr | 0.8.2 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/tinystr/0.8.2/download) |
| tinyvec | 1.11.0 | Zlib OR Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/tinyvec/1.11.0/download) |
| tinyvec_macros | 0.1.1 | MIT OR Apache-2.0 OR Zlib | [source](https://crates.io/api/v1/crates/tinyvec_macros/0.1.1/download) |
| tokio | 1.50.0 | MIT | [source](https://crates.io/api/v1/crates/tokio/1.50.0/download) |
| tokio-macros | 2.6.1 | MIT | [source](https://crates.io/api/v1/crates/tokio-macros/2.6.1/download) |
| tokio-retry | 0.3.1 | MIT | [source](https://crates.io/api/v1/crates/tokio-retry/0.3.1/download) |
| tokio-rustls | 0.26.4 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/tokio-rustls/0.26.4/download) |
| tokio-util | 0.7.18 | MIT | [source](https://crates.io/api/v1/crates/tokio-util/0.7.18/download) |
| tower | 0.5.3 | MIT | [source](https://crates.io/api/v1/crates/tower/0.5.3/download) |
| tower-http | 0.6.8 | MIT | [source](https://crates.io/api/v1/crates/tower-http/0.6.8/download) |
| tower-layer | 0.3.3 | MIT | [source](https://crates.io/api/v1/crates/tower-layer/0.3.3/download) |
| tower-service | 0.3.3 | MIT | [source](https://crates.io/api/v1/crates/tower-service/0.3.3/download) |
| tracing | 0.1.44 | MIT | [source](https://crates.io/api/v1/crates/tracing/0.1.44/download) |
| tracing-appender | 0.2.4 | MIT | [source](https://crates.io/api/v1/crates/tracing-appender/0.2.4/download) |
| tracing-attributes | 0.1.31 | MIT | [source](https://crates.io/api/v1/crates/tracing-attributes/0.1.31/download) |
| tracing-core | 0.1.36 | MIT | [source](https://crates.io/api/v1/crates/tracing-core/0.1.36/download) |
| tracing-log | 0.2.0 | MIT | [source](https://crates.io/api/v1/crates/tracing-log/0.2.0/download) |
| tracing-serde | 0.2.0 | MIT | [source](https://crates.io/api/v1/crates/tracing-serde/0.2.0/download) |
| tracing-subscriber | 0.3.23 | MIT | [source](https://crates.io/api/v1/crates/tracing-subscriber/0.3.23/download) |
| try-lock | 0.2.5 | MIT | [source](https://crates.io/api/v1/crates/try-lock/0.2.5/download) |
| twox-hash | 2.1.2 | MIT | [source](https://crates.io/api/v1/crates/twox-hash/2.1.2/download) |
| typenum | 1.20.1 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/typenum/1.20.1/download) |
| typewit | 1.14.2 | Zlib | [source](https://crates.io/api/v1/crates/typewit/1.14.2/download) |
| unicode-ident | 1.0.24 | (MIT OR Apache-2.0) AND Unicode-3.0 | [source](https://crates.io/api/v1/crates/unicode-ident/1.0.24/download) |
| untrusted | 0.9.0 | ISC | [source](https://crates.io/api/v1/crates/untrusted/0.9.0/download) |
| url | 2.5.8 | MIT OR Apache-2.0 | [source](https://crates.io/api/v1/crates/url/2.5.8/download) |
| urlencoding | 2.1.3 | MIT | [source](https://crates.io/api/v1/crates/urlencoding/2.1.3/download) |
| utf8_iter | 1.0.4 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/utf8_iter/1.0.4/download) |
| uuid | 1.22.0 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/uuid/1.22.0/download) |
| want | 0.3.1 | MIT | [source](https://crates.io/api/v1/crates/want/0.3.1/download) |
| whoami | 2.1.1 | Apache-2.0 OR BSL-1.0 OR MIT | [source](https://crates.io/api/v1/crates/whoami/2.1.1/download) |
| writeable | 0.6.2 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/writeable/0.6.2/download) |
| xet-client | 1.7.0 | Apache-2.0 | [source](https://pypi.org/project/hf-xet/1.7.0/#files) |
| xet-core-structures | 1.7.0 | Apache-2.0 | [source](https://pypi.org/project/hf-xet/1.7.0/#files) |
| xet-data | 1.7.0 | Apache-2.0 | [source](https://pypi.org/project/hf-xet/1.7.0/#files) |
| xet-runtime | 1.7.0 | Apache-2.0 | [source](https://pypi.org/project/hf-xet/1.7.0/#files) |
| yoke | 0.8.1 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/yoke/0.8.1/download) |
| yoke-derive | 0.8.1 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/yoke-derive/0.8.1/download) |
| zerocopy | 0.8.47 | BSD-2-Clause OR Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/zerocopy/0.8.47/download) |
| zerofrom | 0.1.6 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/zerofrom/0.1.6/download) |
| zerofrom-derive | 0.1.6 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/zerofrom-derive/0.1.6/download) |
| zeroize | 1.8.2 | Apache-2.0 OR MIT | [source](https://crates.io/api/v1/crates/zeroize/1.8.2/download) |
| zerotrie | 0.2.3 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/zerotrie/0.2.3/download) |
| zerovec | 0.11.5 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/zerovec/0.11.5/download) |
| zerovec-derive | 0.11.2 | Unicode-3.0 | [source](https://crates.io/api/v1/crates/zerovec-derive/0.11.2/download) |
| zmij | 1.0.21 | MIT | [source](https://crates.io/api/v1/crates/zmij/1.0.21/download) |
