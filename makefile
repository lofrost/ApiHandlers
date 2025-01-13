DB_DSN := "postgres://postgres:loxfrost@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}


migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

migrate-version:
	$(MIGRATE) version

migrate-fix:
	$(MIGRATE) force ${VERSION}

fastgit:
	git add .
	git commit -m "${MESSAGE}"
	git push origin main

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/tasks.gen.go
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/users.gen.go

lint:
	golangci-lint run --out-format=colored-line-number	

run:
	go run cmd/app/main.go