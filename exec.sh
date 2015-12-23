#!/bin/bash

readonly CURRENT_DIR=$(pwd)
readonly SCRIPT_DIR=$(cd $(dirname "$0") && pwd)

cd ${SCRIPT_DIR}/cmd/dochead && go build -v && ./dochead ${SCRIPT_DIR}/parser_test.md && cd ${CURRENT_DIR}
