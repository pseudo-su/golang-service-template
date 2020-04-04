package tools

//go:generate gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=pkg --generate=types -o ../pkg/models.gen.go ../openapi.yaml
////go:generate gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=internal --generate server -o ../internal/server.gen.go ../openapi.yaml
//go:generate gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=pkg --generate=client --make-private=client -o ../pkg/client.gen.go ../openapi.yaml

//go:generate gobin -m -run github.com/pseudo-su/oapi-ui-codegen/cmd/oapi-ui-codegen --package=internal --generate spec -o ../internal/spec.gen.go ../openapi.yaml
//go:generate gobin -m -run github.com/pseudo-su/oapi-ui-codegen/cmd/oapi-ui-codegen --package=internal --generate swaggerui -o ../internal/swagger_ui.gen.go ../openapi.yaml
