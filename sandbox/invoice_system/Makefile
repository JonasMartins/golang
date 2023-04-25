BINARY_TEST = _test
BINARY_DEV = _dev
DONE_MESSAGE = "Finished Operation"

.PHONY: help
## help: shows this help message
help:
	@ echo "Usage: make [target]\n"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
# Docker ###################################################################################

.PHONY: up-dev
## up-dev: build up the related containers
up-dev:
	@ echo "Up dev container"
	@ docker compose -f docker/development/docker-compose.yml --verbose up --build --remove-orphans -d

.PHONY: build-dev
## build-dev: build de executable
build-dev:
	@ rm -f ./${BINARY_DEV} || true && GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o ${BINARY_DEV} ./src/cmd/*.go 
	
.PHONY: down-dev
## down-dev: down the related containers
down-dev:
	@ echo "Down dev container"
	@ docker compose -f docker/development/docker-compose.yml down -v

.PHONY: dev-stop
## dev-stop: stop the related containers
dev-stop:
	@ echo "Stop dev containers"
	@ docker stop development-template-http-mysql development-template-http-mysql-db

.PHONY: dev-start
## dev-start: start the related containers
dev-start:
	@ echo "Start dev containers"
	@ docker start development-template-http-mysql development-template-http-mysql-db

# Local ###################################################################################

.PHONY: test
# test: combine all tests through the entire project
test:
	@ echo "Running general tests..."
	@ rm -f ./${BINARY_TEST} || true && go test ./test/... -c -o ${BINARY_TEST} && ./${BINARY_TEST} -test.v -test.bench=.
	@ echo "Finish running general tests"

.PHONY: dev
## dev: run application 
dev:
	@ echo "Run server"
	@ rm -f ./${BINARY_DEV} || true && GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o ${BINARY_DEV} ./src/cmd/*.go && ./${BINARY_DEV}

.PHONY: clean
## clean: clean binaries
clean:
	@ echo "Removing binaries"
	@ rm -f ./${BINARY_DEV} ./${BINARY_TEST} || true
	@ echo ${DONE_MESSAGE}

# Git ###################################################################################

.PHONY: commit
## commit: add and commit files to develop branch
commit:
	@ git add . && git commit

.PHONY: status
## status: show git status
status:
	@ git status -u

.PHONY: push
## push: push application to main branch
push:
	@ git push -u origin main