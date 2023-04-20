PROJECT_NAME = $(project_name)
PROJECT_PATH = cmd/$(PROJECT_NAME).go

.PHONY:run
run:
	go run $(PROJECT_PATH)

.PHONY:build
build:
	go build -o bin/$(PROGRAM_NAME) $(PROJECT_PATH)

.PHONY:test
test:
	go test ./...

.PHONY:lint
lint:
	golangci-lint run

