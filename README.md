# Golang service template

## Overview

This project contains an example golang template/example application with a number of tooling conventions built in.

For more information on the goals and setup see [./TEMPLATE.md](./TEMPLATE.md)

## Contributing quickstart

### Global dependencies

You will need to have the following global dependencies installed on your local machine

* Go `1.18`
* Sdkman to manage Java versions (only required for PlantUML generation)

```sh
# On macs go can be installed through homebrew
brew install go@1.18

# If using sdkman to manage Java versions
curl -s "https://get.sdkman.io" | bash
sdk env install
sdk env
```

## Local config files

```sh
# Copy the example .env.local file (used to override environment variables for local development)
cp docs/examples/.env.local .

# Copy the recommended vscode settings to your workspace config
cp .vscode/settings.recommended.json .vscode/settings.json
```

If you are using `direnv` (recommended), copy the example `.envrc` file and update it with your desired STAGE name (eg. "stableford"). If you have direnv installed it will automatically load project specific config into your shell when entering the project folder.

```sh
cp docs/examples/.envrc .
# edit `.envrc` file to include your own personal STAGE name
direnv allow
direnv reload
```

## Run development scripts

```sh
# Make sure the correct version of language tooling is active before running any commands
sdk env;

# show available makefile targets
make help;

# Install project dependencies (installs dependencies and tools)
make deps.install;

# Run code verification (static analysis, linting etc)
make verify;

# Verify code using static analysis tools and automatically apply fixes when possible
make verify.fix;

# Run all code generation
make codegen;

# Verify empty git diff after codegen
make verify.empty-git-diff;

# Run unit tests
make test.unit;

# Start local devstack dependencies (Postgres and PgAdmin)
make devstack.start;

# Run whitebox integration tests (needs devstack to be running)
make test.integration.whitebox;

# Execute the database migrations to the latest version
make db.migrate.up.all;

# Run the dev server.
make dev.start;

# Run blackbox integration tests (needs devstack and app to be running)
make test.integration.blackbox;
```

You can stop or recreate the devstack using the following commands

```sh
# Stop shutdown the docker containers running as part of the devstack
make devstack.stop

# Delete/reset the devstack, removes all the containers, volumes etc of the docker-compose stack
make devstack.clean
```
