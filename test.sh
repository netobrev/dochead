#!/bin/bash

readonly CURRENT_DIR=$(pwd)
readonly SCRIPT_DIR=$(cd $(dirname "$0") && pwd)

cd ${SCRIPT_DIR} && go test && cd ${CURRENT_DIR}