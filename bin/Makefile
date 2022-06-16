SERVER_BINARY=serverApp

up:
	@echo "Starting docker images..."
	docker-compose -f docker-compose.staging.yml up
	@echo "Docker images started"

up_build: build_server
	@echo "Stopping docker images (if running...)"
	docker-compose -f docker-compose.staging.yml down
	@echo "Building and starting docker images..."
	docker-compose -f docker-compose.staging.yml up --build -d
	@echo "Docker images built and started."

down:
	@echo "Stopping docker compose..."
	docker-compose -f docker-compose.staging.yml down -v
	@echo "Done"

build_server:
	@echo "Building server binary..."
	cd ../server/src && env GOOS=linux CGO_ENABLED=0 go build -o ${SERVER_BINARY} ./server.go
	@echo "Done"

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

# GIT
status:
	@echo "Git status"
	git status -u

commit:
	@echo "Git add and commit"
	git add . && git commit
	@echo ${DONE}

push:
	@echo "Git push"
	git push -u origin development

reset:
	@echo "Git reset"
	git reset --hard && clean -fd