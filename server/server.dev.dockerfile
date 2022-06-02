FROM golang:1.18-alpine as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o server ./cmd/api
RUN chmod +x /app/server

# build a tiny docker image
FROM alpine:latest  
RUN mkdir /app
COPY --from=builder /app/server /app
CMD ["/app/brokerApp"]


