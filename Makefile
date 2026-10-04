.PHONY: build run

build:
	@go build -o bin/api ./cmd/api/

run: build
	@./bin/api


# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflag="-s -w" -o bin/api ./cmd/api
