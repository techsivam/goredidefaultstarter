# Build stage
FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main

# Final stage
FROM alpine:latest

WORKDIR /

COPY --from=builder /main /main
COPY --from=builder /app/starter.sh /starter.sh

RUN chmod +x /starter.sh

ENV REDIS_HOST=redis
ENV REDIS_PORT=6379
ENV REDIS_PASSWORD=
ENV REDIS_DB=0

CMD ["/starter.sh"]
