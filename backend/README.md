# Backend

This directory contains the Go backend scaffold for Group Events.

Current baseline:
- `cmd/api/main.go` starts the HTTP server and handles graceful shutdown.
- `internal/config/config.go` loads environment-driven runtime config.
- `internal/httpapi/router.go` exposes the initial `GET /health` route.
- `internal/sdui` remains the server-owned UI descriptor seam.
- `internal/domain` defines the first shared backend domain types for groups, memberships, events, polls, votes, and acknowledgements.
- `migrations/` contains SQL migrations managed by goose.
- `internal/testutil` contains shared Postgres-backed test helpers.

Local validation:

```bash
go test ./...
```

Run the command from this `backend/` directory.

## Migrations

Set `DATABASE_URL` and run:

```bash
make migrate-up
make migrate-down
```

## Coverage

Generate backend coverage report:

```bash
make test-cover
```

This writes `coverage.out` in the backend directory.

## Run Locally

```bash
go run ./cmd/api
```

Default API address: `http://localhost:8080`.
