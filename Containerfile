# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.18 as builder

# Copy local code to the container image.
WORKDIR /workdir

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build executable
RUN CGO_ENABLED=0 GOOS=linux go build -v -o service ./cmd/service

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /workdir/service /service

EXPOSE 80/tcp

CMD ["/service"]
