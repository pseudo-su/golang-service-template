export GO111MODULE=on
export GOFLAGS=-mod=vendor

help:
	@echo "This is a helper makefile for oapi-codegen"
	@echo "Targets:"
	@echo "    generate:           rerun code generation"
	@echo "    update-vendor:      tidy and update vendor/"
	@echo "    test-unit:          run unit tests"
	@echo "    test-integration:   run integration tests"
	@echo "    test-smoke:         run smoke tests"

generate:
	go generate ./...

update-vendor:
	go mod tidy
	go mod vendor

test-unit:
	go test ./cmd/... ./internal/... ./pkg/...

test-integration:
	go test ./test-suites/integration/...

test-smoke:
	go test ./test-suites/smoke/...

.PHONY: \
	generate \
	update-vendor \
	test-unit \
	test-integration \
	test-smoke
