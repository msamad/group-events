# group-events

[![Backend CI](https://github.com/msamad/group-events/actions/workflows/backend-ci.yml/badge.svg)](https://github.com/msamad/group-events/actions/workflows/backend-ci.yml)

This repository is a monorepo for the Group Events product.

## Repository layout

- `backend/` contains the Go backend services, API contracts, and database-facing code.
- `mobile/` contains the Flutter mobile app that renders server-driven UI from backend-provided descriptors.
- `packages/` contains shared packages, including SDUI-related building blocks that support the mobile client.
- `.github/` contains automation and workflow definitions.
- `.squad/` contains team operating context, decisions, and agent history.

## How to navigate

Start in the area that matches your task:

- Work on API, data, or backend SDUI descriptors in `backend/`.
- Work on Flutter rendering and mobile experience in `mobile/`.
- Work on delivery automation and feedback loops in `.github/workflows/`.

The top-level folders are scaffolded so CI can target `backend/**` and `mobile/**` independently as the codebase grows.

## Validation

- Backend CI runs on pushes to `main` and pull requests targeting `main` when backend or workflow files change.
- The first workflow validates the Go backend with build, vet, and race-enabled tests.
- Flutter app and package validation are intentionally deferred until both paths are green from a clean baseline.
