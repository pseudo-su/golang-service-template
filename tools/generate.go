package tools

//nolint lll
//go:generate gobin -m -run github.com/pseudo-su/oapi-codegen/cmd/oapi-codegen --package=pkg --generate types -o ../pkg/models.gen.go ../openapi.yaml
//go:generate gobin -m -run github.com/pseudo-su/oapi-codegen/cmd/oapi-codegen --package=internal --generate spec,specui -o ../internal/spec.gen.go ../openapi.yaml
////go:generate gobin -m -run github.com/pseudo-su/oapi-codegen/cmd/oapi-codegen --package=internal --generate server -o ../internal/server.gen.go ../openapi.yaml
//go:generate gobin -m -run github.com/pseudo-su/oapi-codegen/cmd/oapi-codegen --package=pkg --generate client -o ../pkg/client.gen.go ../openapi.yaml
