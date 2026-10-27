# Prefer an image matching the host CPU architecture (amd64 or arm64).
# Cross-architecture emulation may cause model-loading crashes or slow inference.
FROM golang:1.27.3-bookworm@sha256:972b4110ba0f47539a975c31e52c64765e1ff19f2d74469ae4a595c6676e80a5 AS build

WORKDIR /src
# QEMU can hang when Go starts parallel compiler subprocesses.
ENV CGO_ENABLED=1 GOMAXPROCS=1
COPY go.mod go.sum ./
RUN go mod download
COPY Makefile ./
RUN make deps
COPY . .
RUN make build

FROM debian:bookworm-slim@sha256:7ff6ba31527b9e336162fa061746c366c4b2032f4d2cff2078bea43f78efe3f8

# This pinned base already supplies glibc and libstdc++6.
WORKDIR /app
COPY --from=build /src/ez-model-router-bin /usr/local/bin/ez-model-router
COPY --from=build /src/lib/linux-*/lib/libonnxruntime.so.*.* /usr/local/lib/libonnxruntime.so
COPY config.yaml ./config.yaml
COPY LICENSE THIRD_PARTY_NOTICES.md /usr/share/licenses/ez-model-router/
COPY third_party/ /usr/share/licenses/ez-model-router/third_party/

ENV CONFIG_PATH=/app/config.yaml \
    MODEL_DIR=/models \
    ONNX_RUNTIME_PATH=/usr/local/lib/libonnxruntime.so \
    LISTEN_ADDR=:8080
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/ez-model-router"]
