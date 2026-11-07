# Project memory

## Product context
Scam Incident Registry focuses on Nigerian scam reporting and lookup.
Feature scope, acceptance criteria, and progress live in [MVP tickets](docs/mvp-tickets.md).

## Confirmed decisions
- One repository contains a Go backend and React TypeScript frontend using Vite and npm.
- The backend uses standard-library HTTP routing to keep dependencies and abstraction layers small.
- Feature handlers, business rules, and SQL queries belong together in domain packages under backend/internal; add abstractions when real features require them.
- backend/cmd/api is the executable entrypoint; backend/internal/httpapi owns router assembly, health handling, and server lifecycle.
- Shared config, database pools, and S3 clients live under backend/internal/platform; platform packages do not depend on domain packages.
- backend/api/openapi.yaml describes the HTTP contract.
- The frontend uses relative /api URLs through the development proxy.

## Development context
- Toolchain versions are authoritative in backend/go.mod and .nvmrc.
- Backend-only verification: make check-backend build-backend.
- Local infrastructure setup is documented in [infra/README.md](infra/README.md).
- Empty planned directories use .gitkeep until implemented.
- Runtime settings and development commands are documented in [README.md](README.md).
- `make test-integration` loads `backend/.env` when present; `INTEGRATION_ENV_FILE=` selects exported configuration only. Direct Go tests and API startup do not load the file.
- Infrastructure Make targets pass `backend/.env` to Compose when present, including the LocalStack auth token; override with `INFRA_ENV_FILE`. Compose gives exported shell values precedence.

## Outstanding work and decisions
- Domain packages, migrations, and the worker entrypoint remain placeholders.
- The API currently exposes health checks; domain features are tracked separately in the MVP tickets.
- Postgres connections use pgxpool; S3 uses AWS SDK for Go v2. Opt-in client integration tests verify SELECT 1 and bucket access without making API startup depend on infrastructure.
- Domain database queries, background-job execution, production deployment, and operational policies still require decisions.
- Local infrastructure configuration does not imply production service approval or application integration.
- Production hosting must route API requests separately from Vite's development proxy.

## Maintenance
Keep durable decisions and their reasons here. Track feature status in the tickets
and notable delivery changes in CHANGELOG.md. Never store secrets or personal data.
