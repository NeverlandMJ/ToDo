FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o user-service

FROM alpine

COPY --from=builder /app/user-service .
COPY docker.env .env
COPY migrations /app/migrations

CMD ["./user-service"] 