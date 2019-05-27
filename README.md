# Golang Service Template

## Quickstart

### Install Tool Dependencies

```shell
# install golanglint-ci into ./bin
curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.16.0
```

### Build tasks

**NOTE**: if this repo is cloned into you `GOPATH` you will need to prefix all commands with `GO111MODULES=on`

* start: `go run ./...`
* test: `go test ./...`
* lint: `golangci-lint run ./...`

### Build & run docker container

```shell
docker build -t golang-service-template .
docker run -p 8080:80 golang-service-template
```
