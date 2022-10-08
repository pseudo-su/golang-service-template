#!/bin/bash

set -eo pipefail

function errcho(){ >&2 echo $@; }

if [ -z "${CODEGEN_GOJET_DESTINATION}" ]; then
  errcho "ERROR: environment variable CODEGEN_GOJET_DESTINATION not provided"
  exit 1;
fi

if [[ -d "$CODEGEN_GOJET_DESTINATION" ]]; then
  rm -r $CODEGEN_GOJET_DESTINATION
fi

./tools/jet \
  -dsn="postgresql://root:1234@localhost:5432/golang_service_template_localdev?sslmode=disable" \
  -schema=public \
  -ignore-tables="" \
  -path=$CODEGEN_GOJET_DESTINATION

cp -r $CODEGEN_GOJET_DESTINATION/golang_service_template_localdev/public/* $CODEGEN_GOJET_DESTINATION
rm -r $CODEGEN_GOJET_DESTINATION/golang_service_template_localdev
