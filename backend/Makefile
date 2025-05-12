# Makefile

BINARY_NAME=recallify
PORT=8080

# Run the app with Air (live reload)
run:
	air

# Build the Go binary
build:
	go build -o $(BINARY_NAME) ./main.go

# Format Go code
fmt:
	go fmt ./...

# Run tests
test:
	go test ./... -v

# Lint (if using golangci-lint or staticcheck)
lint:
	golangci-lint run

# Migrate DB (example using Bun's migrate tool)
migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

# Docker related (optional if using docker-compose later)
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# Clean up binaries and cache
clean:
	go clean
	rm -f $(BINARY_NAME)
