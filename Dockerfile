FROM alpine:3.18

COPY bin /app/
WORKDIR /app/

CMD ["/app/service"]