# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.12 as builder

# Copy local code to the container image.
WORKDIR /workdir

# Install golangci-lint
RUN curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.16.0

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Run linter and tests
RUN ./bin/golangci-lint run ./...
RUN go test ./...

# Build executable
RUN CGO_ENABLED=0 GOOS=linux go build -v -o service-executable ./internal

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /workdir/service-executable /service-executable

EXPOSE 80/tcp

CMD ["/service-executable"]
