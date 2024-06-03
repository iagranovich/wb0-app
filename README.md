# Help

## Docker

***Run docker compose***
```shell
docker compose -f deploy/docker-compose.yaml --project-directory ./ up -d
```

## Migrations

***Make migration files***
```shell
migrate create -ext sql -dir migrations <migration_mame>
```
***Migration up***
```shell
migrate -path migrations -database postgres://<user:password>@localhost:5432/<db name>?sslmode=disable up
```