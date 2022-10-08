tools/device-info.cfg: tools/device-info.sh
	./tools/device-info.sh > tools/device-info.cfg

tools/golangci-lint: tools/tools.cfg
	. ./tools/tools.cfg && curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./tools v$${golangci_lint}

tools/plantuml.jar: tools/tools.cfg
	. ./tools/tools.cfg && curl -sfL http://sourceforge.net/projects/plantuml/files/plantuml.$${plantuml}.jar/download > ./tools/plantuml.jar

tools/spectral: tools/tools.cfg
	. ./tools/tools.cfg && curl -sfL https://raw.githack.com/stoplightio/spectral/master/scripts/install.sh | sed 's#/usr/local/bin/spectral#./tools/spectral#g' | sh -s -- $${spectral}

tools/migrate: tools/tools.cfg tools/device-info.cfg
	. ./tools/device-info.cfg && . ./tools/tools.cfg && curl -L https://github.com/golang-migrate/migrate/releases/download/v$${golang_migrate}/migrate.$${device_platform}-$${device_architecture}.tar.gz | tar xvz -C ./tools migrate
	touch tools/migrate

tools/conflate: tools/tools.cfg
	. ./tools/tools.cfg && env GOBIN=$${PWD}/tools go install github.com/the4thamigo-uk/conflate/conflate@v$${conflate}

tools/godotenv: tools/tools.cfg
	. ./tools/tools.cfg && env GOBIN=$${PWD}/tools go install github.com/joho/godotenv/cmd/godotenv@v$${godotenv}

tools/oapi-codegen:
	. ./tools/tools.cfg && env GOBIN=$${PWD}/tools go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v$${oapi_codegen}

tools/oapi-ui-codegen:
	. ./tools/tools.cfg && env GOBIN=$${PWD}/tools go install github.com/pseudo-su/oapi-ui-codegen/cmd/oapi-ui-codegen@v$${oapi_ui_codegen}

tools/jet:
	. ./tools/tools.cfg && env GOBIN=$${PWD}/tools go install github.com/go-jet/jet/v2/cmd/jet@v$${gojet}

tools/patter:
	. ./tools/tools.cfg && env GOBIN=$${PWD}/tools go install github.com/apg/patter@$${patter}

tools/mockgen:
	. ./tools/tools.cfg && env GOBIN=$${PWD}/tools go install github.com/golang/mock/mockgen@v$${mockgen}
