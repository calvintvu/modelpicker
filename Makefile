.PHONY: run build clean

run:
	go run .

build:
	go build -o ez-model-router-bin .

clean:
	rm -f ez-model-router-bin
