FROM golang:alpine AS builder

WORKDIR /app

COPY . .

# Build the Go app
RUN go build -o main

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

# Run the main binary when the container starts
CMD [ "./main" ]
