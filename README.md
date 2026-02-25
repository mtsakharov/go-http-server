# go-http-server

A lightweight HTTP/1.1 server built from scratch in Go using raw TCP sockets. No `net/http` — just `net.Conn`, a parser, and a router.

Built as part of the [CodeCrafters "Build Your Own HTTP Server" challenge](https://app.codecrafters.io/courses/http-server/overview).

## Features

- HTTP/1.1 request parsing with keep-alive support
- Request routing (`/`, `/echo/{msg}`, `/files/{name}`, `/user-agent`)
- File serving (GET) and uploading (POST)
- Gzip compression (`Accept-Encoding: gzip`)
- Path traversal protection on file operations
- Concurrent connection handling via goroutines

## Project Structure

```
cmd/server/main.go             # Entry point, TCP listener, connection loop, routing
internal/server/
  httpcore/
    request.go                  # HTTP request parser
    response.go                 # HTTP response writer
    status.go                   # Status code constants
  handlers/
    echo.go                     # /echo/{msg} handler
    files.go                    # /files/{name} handler (GET & POST)
    user_agent.go               # /user-agent handler
  compress/
    gzip.go                     # Gzip compression utility
```

## Requirements

- Go 1.25+

## Usage

### Build and run

```sh
go build -o server ./cmd/server
./server
```

### With a file-serving directory

```sh
./server --directory /tmp/files
```

The server listens on `0.0.0.0:4221`.

### Endpoints

| Method | Path              | Description                          |
|--------|-------------------|--------------------------------------|
| GET    | `/`               | Returns 200 OK                       |
| GET    | `/echo/{message}` | Echoes back the message              |
| GET    | `/user-agent`     | Returns the User-Agent header value  |
| GET    | `/files/{name}`   | Returns file contents from directory |
| POST   | `/files/{name}`   | Writes request body to file          |

### Example

```sh
# Echo
curl http://localhost:4221/echo/hello

# File upload
curl -X POST http://localhost:4221/files/test.txt -d "hello world"

# File download
curl http://localhost:4221/files/test.txt

# Gzip
curl -H "Accept-Encoding: gzip" http://localhost:4221/echo/hello --compressed
```

## Testing

```sh
go test ./...
```

Run with verbose output:

```sh
go test ./... -v
```

## License

MIT
