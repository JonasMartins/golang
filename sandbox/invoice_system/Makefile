#===================##===================##===================##===================##===================##===================##===================##===================#
#====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======##====== ENV ======#
#===================##===================##===================##===================##===================##===================##===================##===================#

BINARY_TEST = _test
BINARY_DEV = _dev
BINARY_BROKER = _broker
BINARY_FINANCIAL = _financial
DONE_MESSAGE = "Finished Operation"

.PHONY: help
## help: shows this help message
help:
	@ echo "Usage: make [target]\n"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: gen-cert
## gen-cert: run generate certificates script
gen-cert:
	@ cd cert && ./gen-dev.sh && cd ..
	@ echo "Task Done"


#===================##===================##===================##===================##===================##===================##===================##===================#
#====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======##====== DOCKER ======#
#===================##===================##===================##===================##===================##===================##===================##===================#

.PHONY: up-dev
## up-dev: build up the related containers
up-dev: build-broker up-broker build-financial up-financial

.PHONY: up-broker
## up-broker: up docker for broker service
up-broker:
	@ echo "Up broker container"
	@ docker compose -f src/services/broker/docker/development/docker-compose.yml --verbose up --build --remove-orphans -d
	@ echo ${DONE_MESSAGE}

.PHONY: up-financial
## up-financial: up docker for financial service
up-financial:
	@ echo "Up financial container"
	@ docker compose -f src/services/financial/docker/development/docker-compose.yml --verbose up --build -d
	@ echo ${DONE_MESSAGE}

.PHONY: build-dev
## build-dev: build de executable
build-dev: build-broker build-financial

.PHONY: build-broker
## build-broker: build de broker service executable
build-broker:
	@ rm -f ./out/${BINARY_BROKER} || true && GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o ./out/${BINARY_BROKER} ./src/services/broker/cmd/*.go 

.PHONY: build-financial
## build-financial: build de broker service executable
build-financial:
	@ rm -f ./out/${BINARY_FINANCIAL} || true && GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o ./out/${BINARY_FINANCIAL} ./src/services/financial/cmd/*.go 

.PHONY: down-dev
## down-dev: down dev containers
down-dev: down-broker down-financial

.PHONY: down-broker
## down-broker: down the broker service related containers
down-broker:
	@ echo "Down broker container"
	@ docker compose -f src/services/broker/docker/development/docker-compose.yml down -v

.PHONY: down-financial
## down-financial: down the financial service related containers
down-financial:
	@ echo "Down financial container"
	@ docker compose -f src/services/financial/docker/development/docker-compose.yml down -v

.PHONY: dev-stop
## dev-stop: stop the related containers
dev-stop:
	@ echo "Stop dev containers"
	@ docker stop micro-go-template-broker-dev micro-go-template-db micro-go-template-financial-dev dev-consul

.PHONY: dev-start
## dev-start: start the related containers
dev-start:
	@ echo "Start dev containers"
	@ docker start micro-go-template-broker-dev micro-go-template-db micro-go-template-financial-dev dev-consul

.PHONY: migrate-create
## migrate-create: run migrations create command usage => make migrate-create migration-name=something
migrate-create:
	@ echo "Running migrations create"
	@ docker compose -f src/services/broker/docker/development/docker-compose.yml --profile tools run --rm migrate create -ext sql -dir /migrations -seq $(migration-name)

.PHONY: migrate-up
## migrate-up: run migrations up command
migrate-up:
	@ echo "Running migrations up"
	@ docker compose -f src/services/broker/docker/development/docker-compose.yml --profile tools run --rm migrate up

.PHONY: migrate-down
## migrate-down: run migrations down command
migrate-down:
	@ echo "Running migrations down"
	@ docker compose -f src/services/broker/docker/development/docker-compose.yml --profile tools run --rm migrate up

.PHONY: migrate-fix
## migrate-fix: run migrations fix command
migrate-fix:
	@ echo "Running migrations fix"
	@ docker compose -f src/services/broker/docker/development/docker-compose.yml --profile tools run --rm migrate force 1


#===================##===================##===================##===================##===================##===================##===================##===================#
#====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======##====== LOCAL ======#
#===================##===================##===================##===================##===================##===================##===================##===================#

.PHONY: test
# test: combine all tests through the entire project
test:
	@ echo "Running general tests..."
	@ rm -f ./out/${BINARY_TEST} || true && go test ./test/... -c -o ./out/${BINARY_TEST} && ./out/${BINARY_TEST} -test.v -test.bench=.
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

.PHONY: proto
## proto: compiles .proto files
proto:
	@ docker run -v ${PWD}:/defs namely/protoc-all -f src/proto/api.proto -l go -o . 

#===================##===================##===================##===================##===================##===================##===================##===================#
#====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======##====== GIT ======#
#===================##===================##===================##===================##===================##===================##===================##===================#

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
