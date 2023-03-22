FROM golang:alpine AS builder

WORKDIR /app

COPY . .

# Build the Go app
RUN go build -o main

FROM alpine:latest

RUN apk update
RUN apk add --no-cache curl mariadb-client

WORKDIR /app

COPY --from=builder /app/main .
COPY init.sh .

EXPOSE 8080

CMD ["sh", "./init.sh"]
