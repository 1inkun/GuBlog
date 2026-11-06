.PHONY: help dev

help:
	@echo "make + [usage]"
	@echo "usage:help,dev"

dev:
	@go vet ./...
	@go fmt ./...
	@go run ./cmd/main/main.go
