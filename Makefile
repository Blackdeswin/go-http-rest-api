.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: butestild
test: 
	go test -v -race -timeout 30s ./...


.DEFAULT_GOAL := build