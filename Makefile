include .env

generate:
	@abigen --abi ./internal/contract/TokenHopper.abi --pkg contract --out ./internal/contract/token_hopper.go
	@go generate ./...

build: generate ## Compile the binary
	@mkdir -p bin
	@go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go

build-linux: generate ## Compile the binary for linux
	@env GOOS=linux go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go

build-docker: build-linux ## Build docker image
	@docker build --build-arg APP_NAME=$(APP_NAME) --build-arg VERSION=$(VERSION) --build-arg ARCH=$(ARCH) -t $(APP_NAME) .

install: build ## compile the binary and copy it to PATH
	@sudo cp build/$(APP_NAME) /usr/local/bin

run: build ## Compile and run the binary
	@./bin/$(APP_NAME)

gomod_tidy: ## Run go mod tidy to clean up & install dependencies
	@go mod tidy

format: ## Run gofumpt against code to format it
	@gofumpt -l -w .

staticcheck: ## Run staticcheck against code
	@staticcheck ./...

test: generate ## Run tests
	@go test -count=1 ./...

codecov-test: generate ## Run tests with coverage
	@mkdir -p coverage
	@go test ./... -coverprofile=coverage/coverage.out
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html

install-deps: install-gofumpt install-mockgen install-abigen install-staticcheck ## Install dependencies

install-gofumpt: ## Install gofumpt for formatting
	go install mvdan.cc/gofumpt@$(GOFUMPT_VERSION)

install-mockgen: ## Install mockgen for generating mocks
	go install github.com/golang/mock/mockgen@$(MOCKGEN_VERSION)

install-staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@$(STATICCHECK_VERSION)

install-abigen: ## install abigen
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest

install-pre-commit: ## Install pre-commit hooks
	@pre-commit install
	@pre-commit install-hooks

pre-commit: ## Run pre-commit hooks
	@pre-commit run --all-files

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
