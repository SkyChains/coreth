#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Lux root directory
CORETH_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )

# Load the versions
source "$CORETH_PATH"/scripts/versions.sh

# Load the constants
source "$CORETH_PATH"/scripts/constants.sh

echo "Building Docker Image: $dockerhub_repo:$build_image_id based of $lux_version"
docker build -t "$dockerhub_repo:$build_image_id" "$CORETH_PATH" -f "$CORETH_PATH/Dockerfile" \
  --build-arg LUX_VERSION="$lux_version" \
  --build-arg CORETH_COMMIT="$coreth_commit" \
  --build-arg CURRENT_BRANCH="$current_branch"
