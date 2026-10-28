# Setup observations

One native run per configuration, 2026-10-28.
Machine: Mac15,6, Apple M3 Pro, 11 logical CPUs, 18 GiB RAM.
Software: macOS 27.0.1, Go 1.27.3, pinned Laya model and native libraries.

| Policy | Launch to first health 200 | RSS at that point |
| --- | --- | --- |
| default config.yaml | 7.03s | 0.61 GiB |
| README example | 1.5s | 0.7 GiB |

Measured from a fresh local checkout using existing model files. Filesystem cache
state was not controlled. Both policies returned valid model selections without fallback.
These are startup snapshots, not peak request memory, minimum RAM, or load benchmarks.

The bundle occupies about 1.6 GiB. Allow 4 GiB for model/temporary transfer files,
plus image/build storage. Native AMD64 startup-time and memory measurements remain pending.
