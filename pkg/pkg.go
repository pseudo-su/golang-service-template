package pkg

//go:generate ../tools/oapi-codegen --package=pkg --generate types -o ./models.gen.go ../openapi/dist/openapi.yml
//go:generate ../tools/oapi-codegen --package=pkg --generate client -o ./client.gen.go ../openapi/dist/openapi.yml
