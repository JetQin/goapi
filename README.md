GO API
======

Overview
--------
Small Go HTTP API demonstrating a minimal service structure using chi router, middleware, and a mock database.

Quick facts
-----------
- Module: github.com/jetqin/goapi
- Entry: `cmd/api/main.go`
- Router: `github.com/go-chi/chi`
- Logger: `github.com/sirupsen/logrus`
- ASCII banner printed with `github.com/common-nighthawk/go-figure`
- Mock database in `internal/tools/mockdb.go`

Repository layout
-----------------
- `cmd/api/main.go` — application entrypoint; sets up router and prints banner.
- `api/api.go` — small HTTP response helpers and error types.
- `internal/handlers/` — HTTP handlers (e.g. `GetCoinBalance`).
- `internal/middleware/` — middleware (authorization, logger).
- `internal/tools/` — mock database and supporting types.

Key components
--------------
- Handlers register routes via `handlers.RegisterHandlers(r)`.
- Middleware chain includes request logging and an authorization check that validates `username` (query) and `Authorization` header.
- `tools.NewDatabaseInterface()` returns a mock `DatabaseInterface` backed by `mockDB` implemented in `internal/tools/mockdb.go`.

Run locally
-----------
1. Ensure dependencies are present (Go >= 1.25 recommended):

```bash
cd /path/to/goapi
go get ./...
```

2. Run the server:

```bash
go run cmd/api/main.go
```

3. Example request (valid credentials present in mock data):

```bash
curl -i 'http://localhost:8080/account/coins?username=alex' -H 'Authorization:123ABC'
```

Expected JSON response:

```json
{"Code":200,"Balance":100}
```

Endpoints
---------
- GET /account/coins?username={username}
  - Headers: `Authorization: <token>`
  - Success: 200 with `CoinBalanceResponse` JSON
  - Errors: 400/500 with `Error` JSON

Troubleshooting
---------------
- Address already in use: another process is listening on port 8080. Find and kill it, or change the port.
  - macOS example: `lsof -i :8080 -sTCP:LISTEN` then `kill <PID>`
- Wrong path: the current handler uses `/account/coins` (plural). If you hit `/account/coin` you'll get 404.
- Missing header/query: Authorization header and `username` query must be present. The middleware returns 400 for missing/invalid credentials.
- Dependencies: run `go get` to ensure third-party modules are installed.

Notes & next steps
------------------
- The project uses an in-memory mock DB for demo; swap `mockDB` for a real DB client by implementing `DatabaseInterface`.
- Consider adding an alias route for `/account/coin` if backward compatibility is required.
- Add unit tests for handlers and middleware.

Contact
-------
For further changes, update the handlers in `internal/handlers` or the mock data in `internal/tools/mockdb.go`.
