FROM golang:1.18.1-alpine3.15
RUN mkdir /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD . /app
RUN go build -o /server

EXPOSE 4001

CMD ["/app/server"]
