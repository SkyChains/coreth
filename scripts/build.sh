#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Root directory
CORETH_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )

# Load the versions
source "$CORETH_PATH"/scripts/versions.sh

# Load the constants
source "$CORETH_PATH"/scripts/constants.sh

if [[ $# -eq 1 ]]; then
    binary_path=$1
elif [[ $# -ne 0 ]]; then
    echo "Invalid arguments to build coreth. Requires either no arguments (default) or one arguments to specify binary location."
    exit 1
fi

# Check if CORETH_COMMIT is set, if not retrieve the last commit from the repo.
# This is used in the Dockerfile to allow a commit hash to be passed in without
# including the .git/ directory within the Docker image.
CORETH_COMMIT=${CORETH_COMMIT:-$(git rev-list -1 HEAD)}

# Build Coreth, which runs as a subprocess
echo "Building Coreth @ GitCommit: $CORETH_COMMIT"
go build -ldflags "-X github.com/skychains/coreth/plugin/evm.GitCommit=$CORETH_COMMIT" -o "$binary_path" "plugin/"*.go
