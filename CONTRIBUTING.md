# Contributing

Contributions are welcome! Here's how to get started.

## Getting Started

1. Fork the repository
2. Clone your fork:
   ```sh
   git clone https://github.com/<your-username>/go-http-server.git
   cd go-http-server
   ```
3. Create a feature branch:
   ```sh
   git checkout -b feature/your-feature
   ```

## Development

### Requirements

- Go 1.25+

### Build

```sh
go build ./cmd/server
```

### Run

```sh
go run ./cmd/server
# or with a file directory:
go run ./cmd/server --directory /tmp/files
```

### Run Tests

```sh
go test ./...
```

Always run tests before submitting a pull request. All tests must pass.

## Project Layout

```
cmd/server/         # Application entry point
internal/server/    # Internal packages (not importable by external modules)
  httpcore/         # Request parsing, response writing, status codes
  handlers/         # Route handlers (echo, files, user-agent)
  compress/         # Gzip compression
```

Key conventions:
- Entry point lives in `cmd/server/main.go`
- All internal packages are under `internal/` to prevent external imports
- Tests live alongside the code they test (`*_test.go`)

## Making Changes

1. Keep changes focused — one feature or fix per PR
2. Add or update tests for any new or changed behavior
3. Run `go vet ./...` and `go test ./...` before committing
4. Write clear commit messages describing *why*, not just *what*

## Pull Requests

1. Push your branch to your fork
2. Open a PR against `master`
3. Describe what the PR does and why
4. Ensure all tests pass

## Code Style

- Follow standard Go conventions (`gofmt`, `go vet`)
- Use meaningful variable and function names
- Keep functions short and focused
- Handle all errors explicitly — no silent `_` on error returns
