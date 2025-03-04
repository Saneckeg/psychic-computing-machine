# Makefile для создания миграций

DB_DSN := "postgres://postgres:4509045@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run cmd/app/main.go

gen-%:
	oapi-codegen -config openapi/.openapi -include-tags $* -package $* openapi/openapi.yaml > ./internal/web/$*/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number