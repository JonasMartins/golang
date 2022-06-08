# ============= RUNNING FIRST TIME AND GENERATE THE BINARY
# FROM golang:1.18-alpine as builder
# RUN mkdir /app
# COPY . /app
# WORKDIR /app
# RUN CGO_ENABLED=0 go build -o serverApp ./src/server.go
# RUN chmod +x /app/server

# # build a tiny docker image
# FROM alpine:latest  
# RUN mkdir /app
# COPY --from=builder /app/server /app
# CMD ["/app/server"]

# ============= RUNNING THE EXECUTABLE AFTER GENERATED FIRST TIME
FROM alpine:latest  

RUN mkdir /app

COPY serverApp /app

CMD ["/app/serverApp"]