# Help

## Docker

***Run docker compose***
```bash
docker compose -f deploy/docker-compose.yaml --project-directory ./ up -d
```

## Migrations

***Make migration files***
```bash
migrate create -ext sql -dir migrations <migration_mame>
```
***Migration up***
```bash
migrate -path migrations -database postgres://<user:password>@localhost:5432/<db name>?sslmode=disable up
```