#!/bin/bash

set -eo pipefail

attempt_counter=0
max_attempts=10

HEALTHCHECK_URI="$1";
HEALTHCHECK_SUCCESS_STATUS_CODE="${2:-200}";

function errcho(){ >&2 echo $@; }

if [ -z "${HEALTHCHECK_URI}" ]; then
  errcho "ERROR: No healthcheck url provided"
  exit 1;
fi

curl_healthcheck () {
  local uri="$1"
  curl --insecure -s -o /dev/null -w ''%{http_code}'' $uri
}

errcho "Waiting for service: uri=$HEALTHCHECK_URI status=$HEALTHCHECK_SUCCESS_STATUS_CODE"

until [[ `curl_healthcheck $HEALTHCHECK_URI` == "$HEALTHCHECK_SUCCESS_STATUS_CODE" ]]; do
    if [ ${attempt_counter} -eq ${max_attempts} ];then
      errcho "Max attempts reached"
      exit 1
    fi

    printf '.'
    attempt_counter=$(($attempt_counter+1))
    sleep 5
done
