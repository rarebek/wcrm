FROM golang:1.20-alpine3.16 AS builder

RUN mkdir app
COPY . /app

WORKDIR /app

RUN go build -o main cmd/app/main.go

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app .

CMD ["/app/main"]