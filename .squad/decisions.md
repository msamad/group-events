# Squad Decisions

## Active Decisions

### 2026-04-20: Server-Driven UI Architecture
**By:** Mustansar Anwar ul Samad (via Danny)
**What:** The Flutter mobile app must implement a server-driven UI architecture. The Go backend returns UI descriptor payloads alongside data responses. The Flutter client renders screens, actions, navigation, and component visibility dynamically based on server responses — no hardcoded screen logic for feature behaviour.
**Why:** Enables updating app behaviour, feature flags, role-based UI, and workflows without requiring app store releases.
**Implications:**
- Backend: All domain API responses must include a `ui` block describing available actions, visible components, and navigation targets for the current user+role context
- Flutter: A core SDUI rendering engine must be built that interprets `ui` descriptors and renders the appropriate widgets
- Test: SDUI contracts must be tested — both the shape of the descriptor from the backend and the rendering behaviour in Flutter
- Future: Notification triggers and in-app prompts can be added server-side without client changes

## Governance

- All meaningful changes require team consensus
- Document architectural decisions here
- Keep history focused on work, decisions focused on direction
