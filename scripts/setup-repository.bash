#!/usr/bin/env bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "$(dirname "${SCRIPT_DIR}/../..")" && pwd)"

cd "${PROJECT_DIR}"

ORIGIN_OWNER="takumin"
ORIGIN_AUTHOR="Takumi Takahashi"
ORIGIN_REPOSITORY="boilerplate-golang-cli"
ORIGIN_DESCRIPTION="Boilerplate Golang CLI Tool"

GITHUB_NAME_WITH_OWNER="$(gh repo view --json nameWithOwner --jq '.nameWithOwner')"

GITHUB_OWNER="${GITHUB_NAME_WITH_OWNER%/*}"
GITHUB_REPOSITORY="${GITHUB_NAME_WITH_OWNER##*/}"
GITHUB_DESCRIPTION="$(gh repo view --json description --jq '.description')"

GITHUB_AUTHOR="$(gh api "users/${GITHUB_OWNER}" --jq '.name')"
[[ -z "${GITHUB_AUTHOR}" ]] && GITHUB_AUTHOR="${GITHUB_OWNER}"

ORIGIN_URL="github.com/${ORIGIN_OWNER}/${ORIGIN_REPOSITORY}"
GITHUB_URL="github.com/${GITHUB_OWNER}/${GITHUB_REPOSITORY}"

go mod edit -module "${GITHUB_URL}"
go-imports-rename -s "${ORIGIN_URL} => ${GITHUB_URL}"

sed -i -E "s@Author.*string.*=.*// ###BOILERPLATE_AUTHOR###@Author string = \"${GITHUB_AUTHOR}\"@" main.go
sed -i -E "s@AppName.*string.*=.*// ###BOILERPLATE_APP_NAME###@AppName string = \"${GITHUB_REPOSITORY}\"@" main.go
sed -i -E "s@AppDesc.*string.*=.*// ###BOILERPLATE_APP_DESC###@AppDesc string = \"${GITHUB_DESCRIPTION}\"@" main.go

gofmt -w .

sed -i -E "s@${ORIGIN_URL}@${GITHUB_URL}@" README.md
sed -i -E "s@${ORIGIN_OWNER}@${GITHUB_OWNER}@" README.md
sed -i -E "s@${ORIGIN_AUTHOR}@${GITHUB_AUTHOR}@" README.md
sed -i -E "s@${ORIGIN_REPOSITORY}@${GITHUB_REPOSITORY}@" README.md
sed -i -E "s@${ORIGIN_DESCRIPTION}@${GITHUB_DESCRIPTION}@" README.md

sed -i -E "s@${ORIGIN_URL}@${GITHUB_URL}@" book.toml
sed -i -E "s@${ORIGIN_OWNER}@${GITHUB_OWNER}@" book.toml
sed -i -E "s@${ORIGIN_AUTHOR}@${GITHUB_AUTHOR}@" book.toml
sed -i -E "s@${ORIGIN_REPOSITORY}@${GITHUB_REPOSITORY}@" book.toml
sed -i -E "s@${ORIGIN_DESCRIPTION}@${GITHUB_DESCRIPTION}@" book.toml
