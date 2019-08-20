# Golang Service Template

## Quickstart

### Install Tool Dependencies

```shell
# install golanglint-ci into ./bin
curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.16.0
# Install gobin globally
env GO111MODULE=off go get -u github.com/myitcv/gobin
```

### Build tasks

**NOTE**: if this repo is cloned into you `GOPATH` you will need to prefix all commands with `GO111MODULES=on`

* start: `go run ./cmd`
* test: `go test ./...`
* lint: `./bin/golangci-lint run ./...`
* build: `go build -o golang-service-template ./cmd`
* code gen: `go generate ./...` [link](https://github.com/go-swagger/go-swagger/issues/1724#issuecomment-469335593)

### Build & run docker container

```shell
docker build -t golang-service-template .
docker run -p 8080:80 golang-service-template
```

## Service Template Details

### Goals

* Encourage reproducible local dev environments by eliminating reliance on globally installed/gopath tools linters etc (`.vscode/settings.json` to set default IDE settings)
* Documentation driven development; Models & API Client generated from openapi v3 spec.
* Embedded documentation (swagger ui) served when deployed to "whitelisted environments".
* Lightweight wrapper for `http.Server` to make bootstrapping the server easy to read and respect proper os process signals for server start/shutdown.
* Unit tests, Integration, and Smoke suite skeletons
* Docker deployment

### Features

* Spec-driven development using `oapi-codegen`
  * `pkg/models.gen.go` Generated Models
  * `internal/spec.gen.go` Embedded swagger spec and swagger ui (exposed at `/swagger.html` and `openapi.json`)
  * `pkg/client.gen.go` Generated API client
* Test suites
  * `test-suites/integration` Integration tests
  * `test-suites/smoke` Smoke tests
* Test reports
  * `reports/test-unit.out`
  * `reports/test-unit.tap`
  * `reports/test-integration.out`
  * `reports/test-integration.tap`
  * `reports/test-smoke.tap`
* When you checkout out the repo for the first time the [`.vscode/settings.json`](https://github.com/pseudo-su/golang-service-template/blob/master/.vscode/settings.json) comes with it (while any further changes you make to it are gitignored).
  * `go.useLanguageServer: true` to use the official language server implementation from google because it’s required in order to support go modules.
  * `go.toolsEnvVars.GO111MODULES: "on"` to make sure that all the go tools know to enable go module support for this project (regardless of if it's inside your gopath or not)
  * `go.toolsEnvVars.GOFLAGS: "-mod=vendor"` if you want to make all go commands use the `vendor/` folder in the project (just like dep did/does)
  * `go.formatTool: "goimports"`: enables automatically managing/formatting package import statements, I think the golang extension uses `gofmt` by default once you enable the `useLanguageServer` flag
  * `go.alternateTools.golangci-lint : "${workspaceFolder}/bin/golangci-lint"` This isn't about go modules, I did it because I want to lock down the version of `golangci-lint` and use a project-local version of it (to avoid issues where people have different versions of golangci-lint globally installed and there’s no way for a project to specify the exact version to use)
