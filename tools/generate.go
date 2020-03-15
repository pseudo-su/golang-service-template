//nolint:go-lint

package tools

// GENERATE TYPES INTO PACKAGES

//go:generate gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=pets --include-tags=pets --generate types -o ../internal/pets/models.gen.go ../openapi.yaml
//go:generate gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=notpets --include-tags=notpets --generate types -o ../internal/notpets/models.gen.go ../openapi.yaml

// GENERATE API CLIENT (pkg)

//go:generate gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=pkg --generate types -o ../pkg/models.gen.go ../openapi.yaml
//go:generate gobin -m -run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --package=pkg --generate client -o ../pkg/client.gen.go ../openapi.yaml

// EMBED SPEC AND OPENAPI

//go:generate gobin -m -run github.com/pseudo-su/oapi-ui-codegen/cmd/oapi-ui-codegen --package=internal --generate spec -o ../internal/spec.gen.go ../openapi.yaml
//go:generate gobin -m -run github.com/pseudo-su/oapi-ui-codegen/cmd/oapi-ui-codegen --package=internal --generate swaggerui -o ../internal/swagger_ui.gen.go ../openapi.yaml
