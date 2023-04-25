## small image
FROM golang:alpine

ARG BINARY_DEV=_dev

WORKDIR /app

COPY src/configs/base.yaml /app/http-mysql/src/configs/
COPY ./$BINARY_DEV /app

CMD ["/app/_dev"]

