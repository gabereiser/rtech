
all: clean deps build
.PHONY: build
build:
	go build
	go build -o bin/rtech cmd/rtech/main.go
	go build -o bin/rcc cmd/asset-compiler/rcc.go
.PHONY: deps
deps:
	go mod tidy
.PHONY: clean
clean:
	go clean
	rm -rf bin;
	rm -rf releases;
	rm -rf packages;

