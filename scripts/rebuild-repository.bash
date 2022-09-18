#!/usr/bin/env bash

set -euo pipefail

if [ ! -x "$(command -v go-imports-rename)" ]; then
	go install github.com/sirkon/go-imports-rename@latest
fi

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "$(dirname "${SCRIPT_DIR}/../..")" && pwd)"

cd "${PROJECT_DIR}"

GITHUB_NAME_WITH_OWNER="$(gh repo view --json nameWithOwner --jq '.nameWithOwner')"
GITHUB_OWNER="${GITHUB_NAME_WITH_OWNER%/*}"
GITHUB_REPOS="${GITHUB_NAME_WITH_OWNER##*/}"

ORIGIN_PKG="github.com/takumin/boilerplate-golang-cli"
PROJECT_PKG="github.com/${GITHUB_OWNER}/${GITHUB_REPOS}"

go mod edit -module "${PROJECT_PKG}"
go-imports-rename -s "${ORIGIN_PKG} => ${PROJECT_PKG}"
