.PHONY: build run

build:
	@go build -o bin/api ./cmd/api/

run: build
	@./bin/api

migrate-up:
	@go run ./cmd/migrate up

migrate-down:
	@go run ./cmd/migrate down


# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflag="-s -w" -o bin/api ./cmd/api
