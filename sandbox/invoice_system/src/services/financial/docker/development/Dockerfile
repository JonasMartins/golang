## small image
FROM alpine:latest

ARG BINARY_DEV=_financial

WORKDIR /app

ADD cert/ /app/invoice_system/cert/
COPY src/services/financial/configs/base.yaml /app/invoice_system/src/services/financial/configs/base.yaml
COPY out/$BINARY_DEV /app

CMD ["/app/_financial"]
