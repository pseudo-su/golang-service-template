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
  -dsn="postgresql://root:1234@localhost:5432/mantel_connect_backend_localdev?sslmode=disable" \
  -schema=public \
  -ignore-tables="allocations,guest_check_ins,schema_migrations,user_ids,check_ins,bookings" \
  -path=$CODEGEN_GOJET_DESTINATION

cp -r $CODEGEN_GOJET_DESTINATION/mantel_connect_backend_localdev/public/* $CODEGEN_GOJET_DESTINATION
rm -r $CODEGEN_GOJET_DESTINATION/mantel_connect_backend_localdev
