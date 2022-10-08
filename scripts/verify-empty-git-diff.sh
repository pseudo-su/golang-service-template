#!/bin/bash

set -eo pipefail

function errcho(){ >&2 echo $@; }

if [[ `git status --porcelain` ]]; then
  errcho "ERROR: uncommited changes detected!";
  git status --porcelain;
  exit 1;
else
  errcho "SUCCESS: No uncommited changes detected";
fi
