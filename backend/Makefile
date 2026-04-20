DATABASE_URL ?= postgres://postgres:postgres@localhost:5432/group_events_test?sslmode=disable

.PHONY: migrate-up migrate-down test test-race test-cover

migrate-up:
	DATABASE_URL=$(DATABASE_URL) go run ./cmd/migrate up

migrate-down:
	DATABASE_URL=$(DATABASE_URL) go run ./cmd/migrate down

test:
	go test ./...

test-race:
	go test -race ./...

test-cover:
	go test ./... -coverprofile=coverage.out
