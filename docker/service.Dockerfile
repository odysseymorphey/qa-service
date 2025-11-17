FROM golang:1.25.0-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o qa-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/qa-service .
COPY ./config ./config
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

ENTRYPOINT [ "./qa-service" ]
