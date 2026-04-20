# Backend

This directory contains the Go backend scaffold for Group Events.

Current baseline:
- `cmd/api/main.go` starts the HTTP server and handles graceful shutdown.
- `internal/config/config.go` loads environment-driven runtime config.
- `internal/httpapi/router.go` exposes the initial `GET /health` route.
- `internal/sdui` remains the server-owned UI descriptor seam.
- `internal/domain` defines the first shared backend domain types for groups, memberships, events, polls, votes, and acknowledgements.

Local validation:

```bash
go test ./...
```

Run the command from this `backend/` directory.
