tidy:
	@echo ">> Tidying..."
	@go mod tidy

fmt:
	@echo ">> Formatting..."
	@go fmt ./...

vet:
	@echo ">> Vetting..."
	@go vet ./...

lint:
	@echo ">> Linting..."
	@revive -config .revive.toml -formatter friendly ./...

sec:
	@echo ">> Auditing..."
	@gosec -quiet ./...

test:
	@echo ">> Running tests..."
	@go test -v -race ./...
.PHONY: test

setup-ci:
	@go get -u github.com/mgechev/revive
	@go get -u github.com/securego/gosec/v2/cmd/gosec

ci: lint vet sec test
