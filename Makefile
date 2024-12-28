# Makefile для создания миграций

# Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down
	
# для удобства добавим команду run, которая будет запускать наше приложение
run:
	go run cmd/app/main.go 

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/task_openapi.yaml > ./internal/web/tasks/taskapi.gen.go
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/user_openapi.yaml > ./internal/web/users/userapi.gen.go

lint:
	golangci-lint run --out-format=colored-line-number