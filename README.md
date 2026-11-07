# Scam Incident Registry

Open-source Nigerian scam incident reporting and lookup system with searchable records and API access.

## Development

- `backend/`: Go HTTP API (Go 1.27.3 or newer).
- `frontend/`: React, TypeScript, and Vite application.

Use Node.js 26.13.0 with npm. With nvm installed:

```sh
nvm install
nvm use
make install
```

Run these in separate terminals from the repository root:

```sh
make dev-backend
make dev-frontend
```

Open the URL printed by Vite (normally http://localhost:5173).
The API listens on port 8080 and exposes `GET /api/health`.
Vite proxies `/api` to that port. To change the backend port, set `PORT`;
update the Vite proxy target to match.

`make check` runs Go tests, Go vet, and frontend lint.
`make build` creates `backend/bin/api` and `frontend/dist/`.
The Vite proxy is development-only; production hosting must route `/api`
to the backend separately.

## Backend structure

```text
backend/
├── cmd/
│   ├── api/
│   └── worker/
├── internal/
│   ├── httpapi/
│   ├── auth/
│   ├── incidents/
│   ├── identifiers/
│   ├── evidence/
│   ├── moderation/
│   ├── disputes/
│   ├── search/
│   ├── apiclients/
│   └── platform/
│       ├── config/
│       ├── database/
│       └── objectstore/
├── migrations/
├── api/
│   └── openapi.yaml
├── .env.example
├── go.mod
└── go.sum
```

`cmd/api` contains the HTTP server entrypoint. `cmd/worker` is reserved for
background jobs, `internal` for domain packages, and `migrations` for database
migrations. These reserved directories contain `.gitkeep` placeholders until
implemented. `api/openapi.yaml` describes the existing HTTP API.

The API uses standard-library routing and structured JSON lifecycle logs.
`internal/platform/config` validates `PORT`; `internal/httpapi` owns router assembly, health handling, and server lifecycle.
Domain packages own feature handlers (`handler.go`), business rules (`service.go`),
and SQL queries (`repository.go`) as those responsibilities are implemented.
`internal/platform/database` manages Postgres pools, and
`internal/platform/objectstore` provides the S3 client. Platform code does not
depend on domain packages.
The server uses 5s header, 15s read, 30s write, and 60s idle timeouts.
SIGINT/SIGTERM allow 10s for requests to drain before forced closure.
Use `make check-backend` and `make build-backend` for Go-only checks and builds.

## Local infrastructure

See [infra/README.md](infra/README.md) for Docker-based Postgres and LocalStack S3,
connection settings, and persistent storage management.

## Project documentation

- [AGENTS.md](AGENTS.md): coding-agent working instructions.
- [MEMORY.md](MEMORY.md): durable context and architecture decisions.
- [MVP tickets](docs/mvp-tickets.md): feature scope, status, and acceptance checklists.
- [CHANGELOG.md](CHANGELOG.md): notable changes by release.

## Check backend connections

Create `backend/.env` from the local-only example if you have not already done so,
and set `LOCALSTACK_AUTH_TOKEN` in that file to your LocalStack token. Infrastructure
commands and integration tests both read this file:

```sh
cp -n backend/.env.example backend/.env
# Set LOCALSTACK_AUTH_TOKEN in backend/.env before starting infrastructure.
make infra-up && make test-integration
```

The opt-in integration tests run `SELECT 1` and S3 `HeadBucket`, with a five-second
timeout per service. They fail if configuration or connectivity is missing and
do not write data. Normal backend checks do not require these services.
Bucket access does not establish upload/download permissions. Configuration is
loaded from `backend/.env` by `make test-integration` when the file exists. File
values override exported variables for this command only. The file must use trusted
shell-compatible assignments. To select another file, use
`make test-integration INTEGRATION_ENV_FILE=backend/.env.example`. To use only
exported variables (for example in CI), use `make test-integration INTEGRATION_ENV_FILE=`.
Direct `go test` commands do not load env files. The API still starts independently
of these services and does not automatically load this file.
For other environments set DATABASE_URL, AWS_REGION, S3_BUCKET, and AWS credentials
through the standard SDK credential chain. S3_ENDPOINT is optional;
S3_USE_PATH_STYLE defaults to false. Never use the example credentials in production.
