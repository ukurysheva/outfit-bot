FROM alpine:3.18

WORKDIR /app

COPY bin ./

CMD ["/app/service"]