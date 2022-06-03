
lint:
	@echo "Running linter"
	cd ../server/src && golangci-lint run
	@echo "Formatting..."
	cd ../server/src && go fmt ./
	@echo "Done"
local:
	@echo "Running development local"
	cd ../server/src && go run server.go
test:
	@echo "Running all tests"
	cd ../server/test && go test ./...