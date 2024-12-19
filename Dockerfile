FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go mod tidy

RUN go build -o user-service cmd/user-service/main.go

CMD ["/app/user-service"]