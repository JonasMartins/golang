# build a tiny docker image
FROM alpine:latest  

RUN mkdir /app

COPY brokerApp /app

CMD ["/app/brokerApp"]

# This create a docker image bigger and build the 
# project, then create another small docker image
# copy the executable and run it on the small image


