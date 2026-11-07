# Repository guidance

## Project
Scam Incident Registry is an open-source Nigerian scam incident
reporting and lookup system.

Read README.md and MEMORY.md before making changes. Consult docs/mvp-tickets.md
for feature context; implement only the requested scope.

## Structure
- backend/: Go API and future background worker.
- frontend/: React, TypeScript, and Vite application.
- infra/: Local development infrastructure.
- docs/: Project planning and documentation.

## Toolchains
- Use the Go version specified in backend/go.mod.
- Use the Node.js version specified in .nvmrc.
- Use npm for frontend dependencies and commit package-lock.json changes.

## Backend conventions
- Keep cmd/api focused on configuration, dependency wiring, and signals.
- Put router assembly, health handling, and server lifecycle in internal/httpapi.
- Put environment loading and validation in internal/platform/config.
- Keep feature handlers, business rules, and SQL queries together in the relevant internal domain package.
- Add handler.go, service.go, and repository.go only as their responsibilities are implemented.
- Keep connection pools in internal/platform/database and S3 integration in internal/platform/objectstore; domain repositories own their queries.
- Platform packages must not depend on domain packages.
- Prefer the Go standard library and explicit dependency injection.
- Add interfaces only where a consuming package needs an abstraction.
- Avoid generic repository layers, global mutable state, and speculative code.
- Propagate request contexts to database and external service calls.
- Return errors to callers; log operational failures at application boundaries.
- Preserve graceful shutdown and server timeouts.
- Update api/openapi.yaml when HTTP contracts change.
- Remove a directory's .gitkeep when adding real files.

## Frontend conventions
- Use TypeScript and functional React components.
- Keep components focused and state close to where it is used.
- Use relative /api URLs; Vite handles the development proxy.
- Handle loading, empty, and error states for asynchronous operations.
- Use semantic HTML and accessible form controls.

## Data and security
- Never commit secrets or real incident evidence.
- Keep .env.example files limited to safe example values.
- Do not log credentials, tokens, evidence contents, or personal identifiers.
- Use synthetic data in tests.
- Add database changes through versioned migrations in backend/migrations.

## Verification
Run checks appropriate to the changed code:
- Backend: make check-backend build-backend
- Frontend: npm --prefix frontend run lint
- Frontend: npm --prefix frontend run build
- All changes: git diff --check

Add focused tests for behavior changes and bug fixes.
Report checks that could not run and explain why.

## Working practices
- Keep changes focused on the requested task.
- Preserve unrelated work in the working tree.
- Add dependencies only when they provide a clear benefit.
- Update documentation when setup, commands, or behavior changes.
- Do not edit generated build outputs or dependency directories.

## Project tracking
- Update the relevant ticket's status and acceptance checklist in docs/mvp-tickets.md when implementing a feature.
- Use Planned, In progress, and Done; mark Done only when all acceptance criteria are implemented and verified.
- Record partial completion and verification evidence without checking off incomplete criteria.
- Keep durable decisions and useful context in MEMORY.md; link to source documentation rather than duplicating it.
- Record notable changes under Unreleased in CHANGELOG.md; do not invent release dates or versions.
