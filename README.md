# Gin WebSocket Boilerplate

Minimal Go/Gin boilerplate that streams `docker ps` output to the browser over WebSocket.

## Features

- WebSocket streaming with gorilla/websocket
- Embedded HTML client (`go:embed`)
- Health check endpoint
- Graceful shutdown (SIGINT/SIGTERM)
- Multi-stage Dockerfile

## Quick Start

```bash
make run
# Open http://localhost:8000
```

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT`   | `8000`  | Server listen port |

## Endpoints

| Method | Path      | Description |
|--------|-----------|-------------|
| GET    | `/`       | WebSocket test client |
| GET    | `/ws`     | WebSocket endpoint (streams `docker ps`) |
| GET    | `/health` | Health check |

## Docker

```bash
make docker-build
make docker-run
```

## Project Structure

```
main.go              # Entrypoint, graceful shutdown
handler/
  websocket.go       # WebSocket handler
  health.go          # Health endpoint
router/
  router.go          # Route definitions
  router_test.go     # Integration tests
web/
  web.go             # go:embed for HTML
  test.html          # WebSocket test client
```

## Development

```bash
make build    # Build binary
make run      # Build and run
make test     # Run tests
make lint     # Run go vet
```

## Author

Baran Gayretli
