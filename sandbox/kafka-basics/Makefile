BINARY_TEST = _test
BINARY_DEV = _dev
DONE_MESSAGE = "Finished Operation"

.PHONY: help
## help: shows this help message
help:
	@ echo "Usage: make [target]\n"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# Local ###################################################################################

.PHONY: clean
## clean: clean binaries
clean:
	@ echo "Removing binaries"
	@ rm -f ./${BINARY_DEV} ./${BINARY_TEST} || true
	@ echo ${DONE_MESSAGE}

.PHONY: build
## build: build application, CGO_ENABLED MUST BE SET TO 1
build:
	@ echo "Build application"
	@ rm -f ./out/${BINARY_DEV} || true && GOOS=linux CGO_ENABLED=1 GOARCH=amd64 go build -o ./out/${BINARY_DEV} ./src/cmd/*.go

.PHONY: produce
## produce: produce some events
produce:
	@ ./out/${BINARY_DEV} project.properties producer
	@ echo ${DONE_MESSAGE}


.PHONY: consume
## consume: consume some events
consume:
	@ ./out/${BINARY_DEV} project.properties consumer
