.PHONY: build
build:
		go build -v cmd/collective-go-sdk/main.go

test:
		go test -v -race -timeout 30s ./ ...

.DEFAULT_GOAL := build