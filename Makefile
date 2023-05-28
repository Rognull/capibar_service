GO ?= go

build: clean ## build application binary
	$(GO) build -v ./cmd/main.go

lint: ## run the linter against the source code
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
	golangci-lint run ./...

clean:
	rm -rf main
