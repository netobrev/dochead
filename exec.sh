#!/bin/bash

readonly CURRENT_DIR=$(pwd)
readonly SCRIPT_DIR=$(cd $(dirname "$0") && pwd)

cd ${SCRIPT_DIR}/cmd/dochead && go build -v && ./dochead -f ${SCRIPT_DIR}/parser_test.ogdl -t ${SCRIPT_DIR}/output_test.tpl  && cd ${CURRENT_DIR}
