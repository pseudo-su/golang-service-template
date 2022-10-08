package internal

//go:generate ../tools/oapi-ui-codegen --package=internal --generate spec -o ./spec.gen.go ../openapi/dist/openapi.yml
//go:generate ../tools/oapi-ui-codegen --package=internal --generate swaggerui -o ./swagger_ui.gen.go ../openapi/dist/openapi.yml
