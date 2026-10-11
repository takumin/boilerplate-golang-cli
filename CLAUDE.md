# Development Guide

## Build Commands
- `task` - Run all tasks (tool, format, generate, reviewdog, test, build)
- `task build` - Build the application for the host (`BUILD_TARGET=<goos>-<goarch>` to cross-compile)
- `task test` - Run all tests with race detection
- `task coverage` - Generate coverage report into `coverage/`
- `task audit` - Scan for known vulnerabilities (govulncheck)
- `task format` - Format code (shfmt, yamlfmt, taplo, pinact, go mod tidy, gofmt)
- `task reviewdog` - Run linters (golangci-lint, govet, gosec, jactionlint, zizmor, etc.)
- Single test: `go test -race -cover ./path/to/package -run TestName`

## Code Style
- Follow Google Go Style Guide conventions ( gofmt )
- Use urfave/cli library for command structure
- Place tests in same directory as implementation files
- Document all public functions
- Use table-driven tests with clear error messages
- Handle errors with context ( fmt.Errorf )
- Use slog for structured logging
- Follow clean architecture with internal packages
- Use Functional Option Pattern for config packages

## Project Structure
- `/internal` - All internal code
- `/internal/command` - CLI command implementation
- `/internal/config` - Configuration management
- `/internal/version` - Version information
- Separate packages by logical concerns

## Error Handling
- Return errors with context
- Use exit codes for command-line errors
- Test error paths with mocked functions

## Pull Requests
- Enable auto-merge right after creating a PR: `gh pr merge <number> --auto --merge`
- If auto-merge is refused because the PR is already mergeable, report it and ask before merging directly
