FROM alpine:latest  

RUN mkdir /app

COPY logerServiceApp /app

CMD ["/app/logerServiceApp"]

