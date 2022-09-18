#!/usr/bin/env bash

set -euo pipefail

if [ ! -x "$(command -v go-imports-rename)" ]; then
	go install github.com/sirkon/go-imports-rename@latest
fi

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "$(dirname "${SCRIPT_DIR}/../..")" && pwd)"

cd "${PROJECT_DIR}"

GITHUB_NAME_WITH_OWNER="$(gh repo view --json nameWithOwner --jq '.nameWithOwner')"
GITHUB_DESCRIPTION="$(gh repo view --json description --jq '.description')"
GITHUB_REPOSITORY="${GITHUB_NAME_WITH_OWNER##*/}"

go mod edit -module "github.com/${GITHUB_NAME_WITH_OWNER}"
go-imports-rename -s "github.com/takumin/boilerplate-golang-cli => github.com/${GITHUB_NAME_WITH_OWNER}"
sed -i -E "s@boilerplate-golang-cli-appname@${GITHUB_REPOSITORY}@" main.go
sed -i -E "s@boilerplate-golang-cli-appdesc@${GITHUB_DESCRIPTION}@" main.go
