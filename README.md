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
