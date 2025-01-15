.PHONY: setup
setup: 
	go mod tidy

.PHONY: run
run: server types
	go run ./cmd/receipt_processor/main.go

# Default target
.PHONY: all
all: server types

.PHONY: server
server:
	oapi-codegen -package=server -generate=server api/api.yml > internal/server/server.gen.go

.PHONY: types
types:
	oapi-codegen -generate types -o internal/server/models/types.gen.go -package models api/api.yml

.PHONY: test
test:
	go test ./... -v 

.PHONY: build
build: server types
	go build -o receipt_processor ./cmd/receipt_processor

.PHONY: clean
clean:
	rm internal/server/server.gen.go
	rm internal/server/models/types.gen.go
	rm receipt_processor

