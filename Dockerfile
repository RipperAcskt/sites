FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./bin/main ./cmd/main.go

EXPOSE 8080

CMD ["./bin/main"]