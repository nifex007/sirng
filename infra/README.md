# Local infrastructure

Requires Docker Compose v2 and a LocalStack auth token with persistence support.
LocalStack persistence uses its authenticated image; obtain a token from your
LocalStack account and put it in the git-ignored `backend/.env` file:

```dotenv
LOCALSTACK_AUTH_TOKEN='your-token'
```

If the file does not exist, copy `backend/.env.example` to `backend/.env` first.
The infrastructure Make targets pass this file to Docker Compose automatically.
An exported shell value takes precedence over the file. Use `INFRA_ENV_FILE=path`
to select another file, or `INFRA_ENV_FILE=` to omit the explicit env file.

```sh
make infra-up
make infra-logs
make infra-down
```

Run these commands from the repository root. Startup waits for both health checks,
including creation of the `sirng-evidence` bucket. Ports are bound to localhost.

## Connections

- Postgres: `postgres://sirng:sirng@localhost:5432/sirng?sslmode=disable`
- S3 endpoint: `http://localhost:4566`
- S3 bucket: `sirng-evidence`
- AWS region: `us-east-1`
- AWS access key and secret: `test` / `test`

Configure your S3 client to use path-style addressing. The AWS test credentials
are separate from the LocalStack auth token. Go client integration tests are available via `make test-integration`; see the root README.

Both services use named volumes. `make infra-down` preserves data. To permanently
remove all local database and S3 data:

```sh
docker compose --env-file backend/.env -f infra/compose.yaml down --volumes
```

See [LocalStack persistence documentation](https://docs.localstack.cloud/aws/developer-tools/snapshots/persistence/).
