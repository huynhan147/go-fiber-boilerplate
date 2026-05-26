APP=myapp

.PHONY: run build tidy test migrate-up migrate-down migrate-reset migrate-version migrate-create

run:
	go run main.go

build:
	go build -o bin/$(APP) main.go

tidy:
	go mod tidy

test:
	go test ./... -v

migrate-create:
	go run cmd/migrate/main.go cmd/migrate/create.go create $(name)

migrate-up:
	go run cmd/migrate/... up

migrate-down:
	go run cmd/migrate/... down

migrate-reset:
	go run cmd/migrate/... reset

migrate-version:
	go run cmd/migrate/... version