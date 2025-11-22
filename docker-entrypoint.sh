#!/bin/sh
set -e

COMMAND=${1:-server}

./bin/$COMMAND "$@"