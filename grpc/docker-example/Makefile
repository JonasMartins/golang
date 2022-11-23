SHELL = /bin/bash

.PHONY: help
## help: shows this help message
help:
	@ echo "Usage: make [target]\n"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: proto
## proto: compiles .proto files
proto:
	@ docker run -v $(PWD):/defs namely/protoc-all -f proto/poetry.proto -l go -o . --go-source-relative

.PHONY: build
## build: builds server's binary
build:
	@ go build -a -installsuffix cgo -o main .

.PHONY: run
## run: runs the server
run: build
	@ ./main

.PHONY: test
## test: runs unit tests
test:
	@ go test -v ./...

.PHONY: build-docker-image
## build-docker-image: builds the docker image
build-docker-image:
	@ docker build . -t docker-grpc-service-tutorial

.PHONY: run-docker
## run-docker: runs the server as a Docker container
run-docker: build-docker-image
	@ docker run -p 4040:4040 docker-grpc-service-tutorial

.PHONY: build-docker-image-multistage
## build-docker-image-multistage: builds a smaller docker image
build-docker-image-multistage:
	@ docker build -f Dockerfile.multistage  . -t docker-grpc-service-tutorial


.PHONY: run-docker-multistage
## run-docker-multistage: runs the server as a Docker container, using the smaller image
run-docker-multistage: build-docker-image-multistage
	@ docker run -p 4040:4040 docker-grpc-service-tutorial
