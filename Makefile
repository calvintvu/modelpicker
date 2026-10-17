.PHONY: run run-inference build clean deps

TOKENIZERS_VERSION := v1.27.0
LIB_DIR := $(CURDIR)/lib
TOKENIZERS_LIB := $(LIB_DIR)/libtokenizers.a

export CGO_LDFLAGS := -L$(LIB_DIR)

$(TOKENIZERS_LIB):
	mkdir -p $(LIB_DIR)
	curl -fsSL https://github.com/daulet/tokenizers/releases/download/$(TOKENIZERS_VERSION)/libtokenizers.linux-amd64.tar.gz \
		| tar -xz -C $(LIB_DIR)

deps: $(TOKENIZERS_LIB)

run: deps
	CONFIG_PATH=./config.yaml \
	MODEL_DIR=./inference/model \
	go run .

run-inference: deps
	go run inference/main.go

build: deps
	go build -o ez-model-router-bin .

clean:
	rm -f ez-model-router-bin


model:
	hf download receptron/laya-onnx \
		--revision 68f27dfe5a27a54fb2b1fefc432f43f972e90868 \
		--local-dir inference/model
