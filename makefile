DB_DSN := "postgres://postgres:loxfrost@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}


migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

fastgit:
	git add .
	git commit -m "fastfix"
	git push origin main

run:
	go run cmd/app/main.go