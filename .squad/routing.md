# Work Routing

How to decide who handles what.

## Routing Table

| Work Type | Route To | Examples |
|-----------|----------|----------|
| Mobile app UX and flows | Rusty | Flutter screens, role-based member actions, acknowledgements, announcements UX |
| Backend API and data contracts | Basher | REST endpoints, auth/roles, groups/memberships, polls and event participation |
| CI and release feedback loop | Livingston | CI pipelines, test feedback loops, integration validation, build automation |
| Code review | Danny | Review PRs, check quality, suggest improvements |
| Testing | Linus | Write tests, run regression checks, verify bug fixes and edge cases |
| Scope & priorities | Danny | What to build next, trade-offs, architecture and milestone decisions |
| Session logging | Scribe | Automatic — never needs routing |

## Issue Routing

| Label | Action | Who |
|-------|--------|-----|
| `squad` | Triage: analyze issue, assign `squad:{member}` label | Lead |
| `squad:danny` | Coordinate architecture, review and route work | Danny |
| `squad:rusty` | Pick up Flutter mobile implementation work | Rusty |
| `squad:basher` | Pick up backend Go/REST implementation work | Basher |
| `squad:linus` | Pick up testing and regression work | Linus |
| `squad:livingston` | Pick up CI, automation and quality pipeline work | Livingston |
| `squad:scribe` | Log sessions and merge decision inbox | Scribe |
| `squad:ralph` | Monitor backlog and maintain work velocity | Ralph |

### How Issue Assignment Works

1. When a GitHub issue gets the `squad` label, the **Lead** triages it — analyzing content, assigning the right `squad:{member}` label, and commenting with triage notes.
2. When a `squad:{member}` label is applied, that member picks up the issue in their next session.
3. Members can reassign by removing their label and adding another member's label.
4. The `squad` label is the "inbox" — untriaged issues waiting for Lead review.

## Rules

1. **Eager by default** — spawn all agents who could usefully start work, including anticipatory downstream work.
2. **Scribe always runs** after substantial work, always as `mode: "background"`. Never blocks.
3. **Quick facts → coordinator answers directly.** Don't spawn an agent for "what port does the server run on?"
4. **When two agents could handle it**, pick the one whose domain is the primary concern.
5. **"Team, ..." → fan-out.** Spawn all relevant agents in parallel as `mode: "background"`.
6. **Anticipate downstream work.** If a feature is being built, spawn the tester to write test cases from requirements simultaneously.
7. **Issue-labeled work** — when a `squad:{member}` label is applied to an issue, route to that member. The Lead handles all `squad` (base label) triage.
