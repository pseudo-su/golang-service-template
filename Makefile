export GO111MODULE=on

help:
	@echo "This is a helper makefile for oapi-codegen"
	@echo "Targets:"
	@echo "    setup:              install the development dependencies"
	@echo "    generate:           rerun code generation"
	@echo "    test-unit:          run unit tests"
	@echo "    test-integration:   run integration tests"
	@echo "    test-smoke:         run smoke tests"

setup:
	# install golanglint-ci into ./bin
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.16.0
	# Install gobin globally
	env GO111MODULE=off go get -u github.com/myitcv/gobin

generate:
	go generate ./...

lint:
	./bin/golangci-lint run ./...

test: test-unit test-integration

report-test: report-test-unit report-test-integration

test-unit:
	go test -count=1 ./cmd/... ./internal/... ./pkg/...

test-integration:
	go test -count=1 ./test-suites/integration/...

test-smoke:
	go test -count=1 ./test-suites/smoke/...

report-test-unit:
	go test -count=1 -coverprofile=reports/test-unit.out -v -p 5 ./cmd/... ./internal/... ./pkg/... | gobin -m -run github.com/apg/patter > reports/test-unit.tap

report-test-integration:
	go test -count=1 -coverprofile=reports/test-integration.out -v -p 5 ./test-suites/integration/... | gobin -m -run github.com/apg/patter > reports/test-integration.tap

report-test-smoke:
	go test -count=1 -coverprofile=reports/test-smoke.out -v -p 5 ./test-suites/smoke/... | gobin -m -run github.com/apg/patter > reports/test-smoke.tap

.PHONY: \
	install-deps \
	generate \
	test-unit \
	test-integration \
	test-smoke \
	report-test-unit \
	report-test-integration \
	report-test-smoke
