export GO111MODULE=on

help:
	@echo "This is a helper makefile for oapi-codegen"
	@echo "Targets:"
	@echo "    setup:              install the development dependencies"
	@echo "    generate:           rerun code generation"
	@echo "    lint:               run linters"
	@echo "    update-deps:        update all go module dependencies"
	@echo "    test-unit:          run unit tests"
	@echo "    test-integration:   run integration tests"
	@echo "    test-smoke:         run smoke tests"
	@echo "    local-ci-build:     emulate a CI build on your local machine"

setup:
	# install golanglint-ci into ./bin
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.24.0
	# install spectral into ./bin
	curl -sfL https://raw.githack.com/stoplightio/spectral/master/scripts/install.sh | sed 's#/usr/local/bin/spectral#./bin/spectral#g' | sed 's#FILENAME=\"spectral\"#FILENAME=\"spectral-linux\"#g' | sh
	# Install gobin globally
	env GO111MODULE=off go get -u github.com/myitcv/gobin
	# Download packages in go.mod file
	go mod download

generate:
	go generate ./...

lint:
	# Lint go files
	./bin/golangci-lint run ./...
	# Lint OpenAPI spec
	./bin/spectral lint --fail-severity=warn ./openapi.yaml

update-deps:
	go get -u ./...
	go mod tidy

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
	cat reports/test-unit.tap

report-test-integration:
	go test -count=1 -coverprofile=reports/test-integration.out -v -p 5 ./test-suites/integration/... | gobin -m -run github.com/apg/patter > reports/test-integration.tap
	cat reports/test-integration.tap

report-test-smoke:
	go test -count=1 -coverprofile=reports/test-smoke.out -v -p 5 ./test-suites/smoke/... | gobin -m -run github.com/apg/patter > reports/test-smoke.tap
	cat reports/test-smoke.tap

run-local-ci-build:
	docker pull nektos/act-environments-ubuntu:18.04
	env CI=true \
	act \
		-P ubuntu-latest=nektos/act-environments-ubuntu:18.04 \
		-j build

.PHONY: \
	help \
	setup \
	generate \
	lint \
	update-deps \
	test-unit \
	test-integration \
	test-smoke \
	report-test-unit \
	report-test-integration \
	report-test-smoke \
	run-local-ci-build
