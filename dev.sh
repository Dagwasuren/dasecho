#!/usr/bin/env bash
set -ex
source .devenv
export HOST="http://dev.dasecho.net:3000"
buffalo db migrate
buffalo dev
