# Perennial Wisdom — Build & Test Automation
# "We suffer more in imagination than in reality." — Seneca

.PHONY: test test-verbose test-cover test-race test-short lint build run clean help

# Default target
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}'

# ---- Build ----

build: ## Build the binary
	go build -o bin/perennial-wisdom .

run: build ## Build and run
	./bin/perennial-wisdom

# ---- Tests ----

test: ## Run all tests
	go test ./... -count=1

test-verbose: ## Run all tests with verbose output
	go test ./... -v -count=1

test-cover: ## Run tests with coverage report
	go test ./... -coverprofile=coverage.out -covermode=atomic -count=1
	go tool cover -func=coverage.out
	@echo ""
	@echo "To view HTML coverage report: make test-cover-html"

test-cover-html: test-cover ## Open coverage report in browser
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

test-race: ## Run tests with race detector
	go test ./... -race -count=1

test-short: ## Run only fast tests (skip integration)
	go test ./... -short -count=1

# ---- Quality ----

lint: ## Run go vet
	go vet ./...

fmt: ## Format code
	gofmt -w .

fmt-check: ## Check formatting (CI-friendly)
	@test -z "$$(gofmt -l .)" || (echo "Files need formatting:" && gofmt -l . && exit 1)

# ---- CI (run all checks) ----

ci: fmt-check lint test-race test-cover ## Run all CI checks

# ---- Clean ----

clean: ## Remove build artifacts
	rm -rf bin/ coverage.out coverage.html
	rm -f wisdom.db
