FROM alpine:3.18

COPY bin/service /app/service
COPY .env /app/.env
WORKDIR /app/

CMD ["/app/service"]