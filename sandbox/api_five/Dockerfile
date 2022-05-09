FROM golang:1.18.1 AS build

WORKDIR /go/src/api_five

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go  build -a -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/api_five/app .

EXPOSE 4001

CMD ["./app"]


