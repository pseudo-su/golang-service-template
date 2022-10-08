# Template

## Project goals

* Encourage reproducible local dev environments by eliminating reliance on globally installed/gopath tools linters etc (`.vscode/settings.json` to set default IDE settings)
* Spec driven development; Models & API Client generated from OpenAPI v3 spec.
* Embedded documentation (swagger ui) served when deployed to "whitelisted environments".
* Lightweight wrapper for `http.Server` to make bootstrapping the server easy to read and respect proper os process signals for server start/shutdown.
* Unit tests, Blackbox integration, and Smoke suite skeletons

## Project features

* Spec-driven development using `oapi-codegen`
  * `pkg/models.gen.go` Generated Models
  * `internal/spec.gen.go` Embedded swagger spec and swagger ui (exposed at `/swagger.html` and `openapi.json`)
  * `pkg/client.gen.go` Generated API client
* Test suites
  * `test_suites/integration_whitebox` Whitebox integration tests
  * `test_suites/integration_blackbox` Blackbox integration tests
  * `test_suites/smoke` Smoke tests
* Test reports
  * `reports/test-unit.out`
  * `reports/test-unit.tap`
  * `reports/test-integration.out`
  * `reports/test-integration.tap`
  * `reports/test-smoke.tap`
