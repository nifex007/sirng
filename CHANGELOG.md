# Changelog

Notable project changes are recorded here. Unreleased entries describe work in
the repository, not a published release.

## Unreleased

### Fixed
- Infrastructure Make targets also read `backend/.env`, allowing the LocalStack auth token to be configured there.
- Integration-test Make target loads local `backend/.env` configuration, with an explicit env-file override and an environment-only option for CI.

### Added
- Reusable Postgres and S3 clients, local configuration examples, and opt-in client integration tests.
- Go backend and React TypeScript/Vite frontend in a monorepo, with pinned toolchains and development commands.
- Backend domain directory structure and OpenAPI health endpoint specification.
- Validated API configuration, structured lifecycle logging, HTTP timeouts, graceful shutdown, and focused tests.
- Local Postgres and LocalStack S3 infrastructure configuration and setup documentation.
- Project-specific ignore rules and contributor guidance.
- Project memory, MVP ticket statuses and acceptance checklists, and this changelog.

### Changed
- Grouped backend configuration, Postgres pools, and S3 clients under internal/platform; established domain-owned feature handlers and repositories.
