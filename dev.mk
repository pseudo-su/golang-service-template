include tools/tools.mk

### Dependencies - manage project and tool dependencies

## Install dependencies
deps.install: deps.tools.install deps.app.install
.PHONY: deps.install

## Update dependencies
deps.update: deps.tools.update deps.app.update
.PHONY: deps.update

## Install app dependencies
deps.app.install:
	# Golang dependencies (go.mod, go.sum)
	go mod download

	touch .env.local
	touch test_suites/.env.local
	touch test_suites/.env.dev.local
.PHONY: deps.app.install

## Update app dependencies
deps.app.update:
	# Golang dependencies (go.mod, go.sum)
	go get -u ./...
	go mod tidy
.PHONY: deps.app.update

## Install tool dependencies
deps.tools.install: tools/golangci-lint tools/plantuml.jar tools/spectral tools/migrate tools/conflate tools/godotenv tools/oapi-codegen tools/oapi-ui-codegen tools/jet tools/patter
.PHONY: deps.tools.install

## Update tool dependencies
deps.tools.update: deps.tools.install
	>&2 echo "WARNING: Any tool dependencies need to be updated manually"
.PHONY: deps.tools.update

### Test

TEST_DOTENV := ./tools/godotenv -f ./test_suites/.env.local,./test_suites/.env
TEST_DOTENV_STAGE := ./tools/godotenv -f ./test_suites/.env.${STAGE}.local,./test_suites/.env.${STAGE},./test_suites/.env.local,./test_suites/.env
TEST_GO := go test -count=1 -v -p=1
TEST_GO_REPORT := bash -c 'tee >(cat 1>&2)' | ./tools/patter

precheck-report:
	mkdir -p reports
.PHONY: precheck-report

test.precheck-dotenv:
	touch test_suites/.env.local
	if [ ! -z "${STAGE}" ]; then \
		touch test_suites/.env.${STAGE}.local; \
	fi
.PHONY: test.precheck-dotenv

## Run unit tests
test.unit: test.unit.go test.unit.openapi
.PHONY: test.unit

## Run unit tests and report
test.unit.report: test.unit.go.report test.unit.openapi
.PHONY: test.unit.report

## Run Go unit tests
test.unit.go:
	${TEST_GO} ./cmd/... ./internal/... ./pkg/...
.PHONY: test.unit.go

## Run Go unit tests and generate reports
test.unit.go.report: precheck-report
	${TEST_GO} -coverprofile=reports/test.unit.go.out ./cmd/... ./internal/... ./pkg/... | ${TEST_GO_REPORT} > reports/test.unit.go.tap
.PHONY: test.unit.go.report

## Run OpenAPI unit tests
test.unit.openapi:
	echo "TODO: OpenAPI unit tests"
	# ${TEST_DOTENV} ${TEST_GO} ./test_suites/openapi/...
.PHONY: test.unit.openapi

## Run all integration tests
test.integration: test.integration.whitebox test.integration.blackbox
.PHONY: test.integration

## Run all integration tests and generate reports
test.integration.report: test.integration.whitebox.report test.integration.blackbox.report
.PHONY: test.integration.report

## Run whitebox integration tests (require devstack to be running)
test.integration.whitebox: test.precheck-dotenv
	${TEST_DOTENV} ${TEST_GO} ./test_suites/integration_whitebox/...
.PHONY: test.integration.whitebox

## Run whitebox integration tests and generate reports (require devstack to be running)
test.integration.whitebox.report: precheck-report test.precheck-dotenv
	${TEST_DOTENV} ${TEST_GO} ./test_suites/integration_whitebox/... | ${TEST_GO_REPORT} > reports/test.integration.whitebox.tap
.PHONY: test.integration.whitebox.report

## Run blackbox integration tests (require devstack and app to be running)
test.integration.blackbox: test.precheck-dotenv
	${TEST_DOTENV} ${TEST_GO} ./test_suites/integration_blackbox/...
.PHONY: test.integration.blackbox

## Run blackbox integration tests and generate reports (require devstack and app to be running)
test.integration.blackbox.report: precheck-report test.precheck-dotenv
	${TEST_DOTENV} ${TEST_GO} ./test_suites/integration_blackbox/... | ${TEST_GO_REPORT} > reports/test.integration.blackbox.tap
.PHONY: test.integration.blackbox.report

## Run smoke tests
test.smoke: test.precheck-dotenv
	if [ -z "${STAGE}" ]; then echo "STAGE environment variable not set"; exit 1; fi
	${TEST_DOTENV_STAGE} ${TEST_GO} ./test_suites/smoke/...
.PHONY: test.smoke

## Run smoke tests and generate reports
test.smoke.report: test.precheck-dotenv
	if [ -z "${STAGE}" ]; then echo "STAGE environment variable not set"; exit 1; fi
	${TEST_DOTENV_STAGE} ${TEST_GO} ./test_suites/smoke/... | ${TEST_GO_REPORT} > reports/test.smoke.${STAGE}.tap
.PHONY: test.smoke.report

### Verify - Code verifiation and Static analysis

## Run code verification
verify: verify.go verify.openapi
.PHONY: verify

## Run code verifiation and automatically apply fixes where possible
verify.fix: verify.go.fix verify.openapi
.PHONY: verify.fix

# Run static analysis on Golang code
verify.go: precheck-report
	./tools/golangci-lint run --timeout 3m0s ./...
.PHONY: verify.go

# Run static analysis on Golang code and autofix where possible
verify.go.fix: precheck-report
	./tools/golangci-lint run --timeout 3m0s --fix ./...
.PHONY: verify.go.fix

# Run static analysis on OpenAPI spec
verify.openapi:
	./tools/spectral lint --ruleset=.spectral.yml -v --fail-severity=warn ./openapi.yml
.PHONY: verify.openapi

## Verify empty commit diff after codegen
verify.empty-git-diff:
	./scripts/verify-empty-git-diff.sh
.PHONY: verify.empty-git-diff

### Code generation

## Run all code generation
codegen: codegen.docs codegen.openapi codegen.go
.PHONY: codegen

## Run docs code generation
codegen.docs:
	./scripts/generate-docs.sh
.PHONY: codegen.docs

## Run OpenAPI code generation
codegen.openapi:
	./tools/conflate -data ./openapi/_template.yml -format YAML > ./openapi/dist/openapi.yml
.PHONY: codegen.openapi

## Run Golang code generation
codegen.go:
	go generate ./...
.PHONY: codegen.go

## Run Gojet codegen (requires devstack database)
codegen.gojet:
	env CODEGEN_GOJET_DESTINATION=./internal/persistence/gojet ./scripts/generate-gojet-code.sh
.PHONY: codegen.gojet

### Dev

DEV_DOCKER_COMPOSE := docker-compose -f ./devstack/app.docker-compose.yml -p app

## Run the development server
dev.run:
	go run ./cmd/service
.PHONY: dev.run

## Start the development server in the background
dev.start:
	 ${DEV_DOCKER_COMPOSE} up -d --remove-orphans app
.PHONY: dev.start

## Stop the development server
dev.stop:
	${DEV_DOCKER_COMPOSE} down --remove-orphans
.PHONY: dev.stop

# Clean/reset the development server
dev.clean: dev.stop
	${DEV_DOCKER_COMPOSE} down --remove-orphans --volumes --rmi local
.PHONY: dev.clean

## Restart  the development server
dev.restart: dev.stop dev.start
.PHONY: dev.restart

## Recreate and restart the development server
dev.recreate: dev.clean dev.start
.PHONY: dev.recreate

## Show the status of the development server
dev.status:
	${DEV_DOCKER_COMPOSE} ps
.PHONY: dev.status

## Show the logs of the development server
dev.logs:
	${DEV_DOCKER_COMPOSE} logs --follow
.PHONY: dev.logs

## Export logs to reports folder
dev.logs.report: precheck-report
	${DEV_DOCKER_COMPOSE} logs | bash -c 'tee >(cat 1>&2)' > reports/dev.log
.PHONY: dev.logs.report


### Devstack

DEVSTACK_DOCKER_COMPOSE := docker-compose -f ./devstack/devstack.docker-compose.yml -p devstack

## Start the devstack
devstack.start:
	${DEVSTACK_DOCKER_COMPOSE} up -d --remove-orphans devstack
.PHONY: devstack.start

## Stop the devstack
devstack.stop:
	${DEVSTACK_DOCKER_COMPOSE} down --remove-orphans
.PHONY: devstack.stop

## Clean/reset the devstack
devstack.clean:
	${DEVSTACK_DOCKER_COMPOSE} down --remove-orphans --volumes --rmi local
.PHONY: devstack.clean

## Restart the devstack
devstack.restart: devstack.stop devstack.start
.PHONY: devstack.restart

## Clean/reset and restart the devstack
devstack.recreate: devstack.clean devstack.start
.PHONY: devstack.recreate

## Show status
devstack.status:
	${DEVSTACK_DOCKER_COMPOSE} ps
.PHONY: devstack.status

## Show logs
devstack.logs:
	${DEVSTACK_DOCKER_COMPOSE} logs --follow
.PHONY: devstack.logs

## Export logs to reports folder
devstack.logs.report: precheck-report
	${DEVSTACK_DOCKER_COMPOSE} logs | bash -c 'tee >(cat 1>&2)' > reports/devstack.log
.PHONY: devstack.logs.report

### Database management

DB_MIGRATE := ./tools/migrate -database "postgres://root:1234@localhost:5432/mantel_connect_backend_localdev?sslmode=disable" -source file://./internal/persistence/migrations

## Create a new migration in ./internal/persistence/migrations folder
db.migrate.create:
	if [ -z "${NEW_MIGRATION_NAME}" ]; then echo "NEW_MIGRATION_NAME environment variable not set"; exit 1; fi
	${DB_MIGRATE} create -dir ./internal/persistence/migrations -ext sql $${NEW_MIGRATION_NAME}
.PHONY: db.migrate.create

## Print current migration version
db.migrate.version:
	${DB_MIGRATE} version
.PHONY: db.migrate.version

## Run all migrations up to the latest
db.migrate.up.all:
	${DB_MIGRATE} up
.PHONY: db.migrate.up.all

## Run the next up migration
db.migrate.up.1:
	${DB_MIGRATE} up 1
.PHONY: db.migrate.up.1

## Revert all migrations down to the first
db.migrate.down.all:
	${DB_MIGRATE} down
.PHONY: db.migrate.down.all

## Revert all migrations down to the first
db.migrate.down.1:
	${DB_MIGRATE} down 1
.PHONY: db.migrate.down.1
