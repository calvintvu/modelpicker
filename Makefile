.PHONY: deps check-platform run build clean model
.DELETE_ON_ERROR:

TOKENIZERS_VERSION := v1.27.0
ONNXRUNTIME_VERSION := 1.29.0
MODEL_DIR ?= inference/model
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
PLATFORM := $(GOOS)-$(GOARCH)
LIB_DIR ?= $(CURDIR)/lib/$(PLATFORM)
TOKENIZERS_LIB := $(LIB_DIR)/libtokenizers.a
RUNTIME_LIBRARY := $(if $(filter darwin,$(GOOS)),libonnxruntime.dylib,libonnxruntime.so)
ONNXRUNTIME_LIB := $(LIB_DIR)/lib/$(RUNTIME_LIBRARY)

ONNXRUNTIME_PLATFORM_linux-amd64 := linux-x64
ONNXRUNTIME_PLATFORM_linux-arm64 := linux-aarch64
ONNXRUNTIME_PLATFORM_darwin-arm64 := osx-arm64
ONNXRUNTIME_PLATFORM := $(ONNXRUNTIME_PLATFORM_$(PLATFORM))

TOKENIZERS_SHA256_linux-amd64 := 72556cdca798dd4ea7cdaba308e5f0d68a8cb93b67c96edf485b7a0edd7b07f4
TOKENIZERS_SHA256_linux-arm64 := e96545ad05930c26f51f63d932ee6d3bbd32bbed149e102c5290d587a2293067
TOKENIZERS_SHA256_darwin-arm64 := fb84b8b2e349a5952767ffe80ccd862fc44084de47f3b0cc3f0b7c9d4e649cf7
TOKENIZERS_SHA256 := $(TOKENIZERS_SHA256_$(PLATFORM))
ONNXRUNTIME_SHA256_linux-amd64 := c3fddc4f139a045b0c4902c57410f0694f1c2fdf9b6939fbe38b1aeae7cd14ba
ONNXRUNTIME_SHA256_linux-arm64 := e1799098ebc054b370f6176a450f158720f297818c613e5dc99b92e2ec82346f
ONNXRUNTIME_SHA256_darwin-arm64 := d0706fc34f315d8c88639d0a8c81f2e09e815f282cabed3493c06a054352cf92
ONNXRUNTIME_SHA256 := $(ONNXRUNTIME_SHA256_$(PLATFORM))

export CGO_LDFLAGS := $(CGO_LDFLAGS) -L$(LIB_DIR)

deps: $(TOKENIZERS_LIB) $(ONNXRUNTIME_LIB)

check-platform:
	@test -n "$(ONNXRUNTIME_PLATFORM)" || { echo "Native dependencies are available for linux-amd64, linux-arm64, and darwin-arm64; $(PLATFORM) is deferred." >&2; exit 1; }

$(TOKENIZERS_LIB) $(ONNXRUNTIME_LIB): | check-platform

$(TOKENIZERS_LIB):
	mkdir -p "$(LIB_DIR)"
	curl -fsSL "https://github.com/daulet/tokenizers/releases/download/$(TOKENIZERS_VERSION)/libtokenizers.$(PLATFORM).tar.gz" -o "$@.tar.gz"
	printf '%s  %s\n' "$(TOKENIZERS_SHA256)" "$@.tar.gz" | shasum -a 256 -c -
	tar -xzf "$@.tar.gz" -C "$(LIB_DIR)"
	rm "$@.tar.gz"

# The macOS archive has an extra leading ./ component.
$(ONNXRUNTIME_LIB):
	mkdir -p "$(LIB_DIR)/lib"
	curl -fsSL "https://github.com/microsoft/onnxruntime/releases/download/v$(ONNXRUNTIME_VERSION)/onnxruntime-$(ONNXRUNTIME_PLATFORM)-$(ONNXRUNTIME_VERSION).tgz" -o "$@.tar.gz"
	printf '%s  %s\n' "$(ONNXRUNTIME_SHA256)" "$@.tar.gz" | shasum -a 256 -c -
	tar -xzf "$@.tar.gz" --strip-components=$(if $(filter darwin,$(GOOS)),2,1) -C "$(LIB_DIR)"
	rm "$@.tar.gz"

run: deps
	go run .

build: deps
	go build -o ez-model-router-bin .

clean:
	rm -f ez-model-router-bin

model:
	@command -v hf >/dev/null 2>&1 || { echo "Hugging Face CLI (hf) not found. Install with: uv tool install huggingface_hub" >&2; exit 1; }
	hf download receptron/laya-onnx \
		--revision 68f27dfe5a27a54fb2b1fefc432f43f972e90868 \
		--local-dir "$(MODEL_DIR)"
