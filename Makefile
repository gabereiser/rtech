

build:
	go build
	go build -o bin/rtech cmd/rtech/main.go
	go build -o bin/rcc cmd/asset-compiler/rcc.go
deps:
	go mod tidy
clean:
	go clean
	rm -rf bin;
	rm -rf releases;
	rm -rf packages;
all: clean deps build
