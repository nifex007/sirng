.PHONY: install dev-backend dev-frontend build check
install:
	cd frontend && npm ci
dev-backend:
	cd backend && go run ./cmd/api
dev-frontend:
	cd frontend && npm run dev
build:
	cd backend && go build -o bin/api ./cmd/api
	cd frontend && npm run build
check:
	cd backend && go test ./... && go vet ./...
	cd frontend && npm run lint

.PHONY: check-backend build-backend
check-backend:
	cd backend && go vet ./... && go test -race ./...
	cd backend && test -z "$$(gofmt -l cmd internal)"
build-backend:
	cd backend && go build -o bin/api ./cmd/api

.PHONY: infra-up infra-down infra-logs
INFRA_ENV_FILE ?= $(wildcard backend/.env)
COMPOSE = docker compose $(if $(INFRA_ENV_FILE),--env-file "$(INFRA_ENV_FILE)") -f infra/compose.yaml
infra-up:
	$(COMPOSE) up -d --wait
infra-down:
	$(COMPOSE) down
infra-logs:
	$(COMPOSE) logs -f

.PHONY: test-integration
INTEGRATION_ENV_FILE ?= $(wildcard backend/.env)
test-integration:
	@set -e; \
	if [ -n "$(INTEGRATION_ENV_FILE)" ]; then \
		set -a; . "$(abspath $(INTEGRATION_ENV_FILE))"; set +a; \
	fi; \
	cd backend; \
	go test -race -tags=integration -count=1 -timeout=30s ./internal/platform/database ./internal/platform/objectstore
