#!/bin/bash

set -eu

THIS_FILE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
CI="${THIS_FILE_DIR}/../../wg-app-platform-runtime-ci"
. "$CI/shared/helpers/git-helpers.bash"
REPO_NAME=$(git_get_remote_name)

CONTAINER_NAME="$REPO_NAME-docker-container"

DB="${DB:-mysql-8.0}" "${THIS_FILE_DIR}/create-docker-container.bash" -d

docker exec \
  --env DB="${DB:-mysql-8.0}" \
  $CONTAINER_NAME \
  '/repo/scripts/docker/test.bash' "$@"
