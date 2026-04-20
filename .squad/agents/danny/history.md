# Project Context

- **Owner:** Mustansar Anwar ul Samad
- **Project:** group-events
- **Description:** Flutter mobile app + Go backend with REST APIs for event scheduling, role-based participation, polls, and acknowledgements across multi-group memberships.
- **Stack:** Flutter, Go, REST API, CI automation, regression testing
- **Created:** 2026-04-20

## Learnings

Initial team cast selected from Ocean's Eleven universe.
Testing and regression confidence are core delivery requirements.

## Sprint 1 Planning Complete — 2026-04-20

Architecture directive received: Server-Driven UI (SDUI).
- Recorded SDUI decision to decisions inbox
- Updated issues #2, #3, #5, #7, #8 with SDUI alignment notes
- Created 3 new SDUI-specific issues: descriptor types + rendering engine, contract tests, role-aware actions
- Sprint 1 board: 15 issues total covering scaffold, domain models, SDUI core, Groups CRUD, test scaffolds, CI pipeline
- All Flutter feature issues must use SduiRenderer — no hardcoded role logic in client
