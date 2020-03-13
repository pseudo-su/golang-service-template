export GO111MODULE=on

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

test: test-unit test-integration

report-test: report-test-unit report-test-integration

test-unit:
	go test ./cmd/... ./internal/... ./pkg/...

test-integration:
	go test ./test-suites/integration/...

test-smoke:
	go test ./test-suites/smoke/...

report-test-unit:
	go test -coverprofile=reports/test-unit.out -v -p 5 ./cmd/... ./internal/... ./pkg/... | gobin -m -run github.com/apg/patter > reports/test-unit.tap

report-test-integration:
	go test -coverprofile=reports/test-integration.out -v -p 5 ./test-suites/integration/... | gobin -m -run github.com/apg/patter > reports/test-integration.tap

report-test-smoke:
	go test -v -p 5 ./test-suites/smoke/... | gobin -m -run github.com/apg/patter > reports/test-smoke.tap

.PHONY: \
	generate \
	update-vendor \
	test-unit \
	report-test-unit \
	test-integration \
	report-test-integration \
	test-smoke \
	report-test-smoke
