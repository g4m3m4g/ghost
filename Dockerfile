# Build stage
FROM golang:1.26.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ghost ./server

# Run stage (lightweight)
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/ghost .

EXPOSE 8080

CMD ["./ghost"]