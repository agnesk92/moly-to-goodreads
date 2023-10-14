makefile := $(firstword $(MAKEFILE_LIST))

##
## 🔨 GENERAL
##

.PHONY: help
help:  ## Help.
	@echo "Targets:"
	@grep -F -h "##" $(makefile) | grep -F -v grep -F | sed -e 's/\\$$//' | sed -e 's/\(.*\):.*##[ \t]*/    \1 ## /' | column -t -s '##'
	@echo

.PHONY: install
install:  ## Install utilities.
	@echo "Installing dependencies..."
	@go install golang.org/x/tools/cmd/goimports@latest
	@@$(shell curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin latest)

##
## ✨ TESTING & QUALITY
##

.PHONY: test
test: ## Run unit tests.
	@echo "Running tests"
	@go test $$(go list ./...)

.PHONY: fmt
fmt: ## Format the code.
	@echo "Formatting the code"
	@gofmt -s -w .

.PHONY: lint
lint:  ## Lint code.
	@echo "Linting the code"
	@golangci-lint -v --color always run --sort-results ./...
