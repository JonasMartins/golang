## small image
FROM alpine:latest

ARG BINARY_DEV=_broker

WORKDIR /app

RUN echo "$PWD"

COPY src/services/broker/configs/base.yaml /app/invoice_system/src/services/broker/configs/base.yaml
COPY out/$BINARY_DEV /app

CMD ["/app/_broker"]



# FROM golang:1.20-alpine3.17 as builder

# ARG BINARY_DEV=_broker

# # removing unecessaries binaries
# RUN rf -f ./out/$BINARY_DEV || true

# WORKDIR /app

# COPY ../../../../../go.mod ./
# COPY ../../../../../go.sum ./

# RUN go mod download && go mod verify

# COPY ../../ .

# RUN go build -o $BINARY_DEV ./src/cmd/*.go

# RUN chmod +x ./$BINARY_DEV

# ## small image
# FROM golang:alpine

# WORKDIR /app

# RUN mkdir -p app/http-mysql/src/cmd/configs

# COPY src/configs/base.yaml /app/http-mysql/src/configs/
# COPY --from=builder $BINARY_DEV /app

# CMD ["/app/_dev"]
