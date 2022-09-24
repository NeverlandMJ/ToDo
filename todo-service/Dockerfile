FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o todo-service

FROM alpine

COPY --from=builder /app/todo-service .
COPY docker.env .env
COPY migrations /app/migrations

CMD ["./todo-service"] 