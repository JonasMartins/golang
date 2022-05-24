# get the base go image
FROM golang:1.18-alpine as builder

# create a dir on docker image
RUN mkdir /app

# copy everuthing from current folder into 
# recent created app folder
COPY . /app

# set working directory
WORKDIR /app

# add a environment variable
# and run the application 
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# make sure the exe file os executable
RUN chmod +x /app/brokerApp

# build a tiny docker image
FROM alpine:latest  

RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]

# This create a docker image bigger and build the 
# project, then create another small docker image
# copy the executable and run it on the small image


