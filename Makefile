GOFILES := $(shell find . -name "*.go" -type f ! -path "./vendor/*")
GOFMT ?= gofmt -s

.PHONY: build
build:
	CGO_ENABLED=0 go build -a -o ./bin/hatch ./main/

.PHONY: test
test:
	go test -v -race ./main/... ./pkg/...