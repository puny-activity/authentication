FROM golang:1.23.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app cmd/httpserver/main.go

CMD ["./app"]
