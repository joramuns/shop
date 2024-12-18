FROM golang:1.23 as builder

WORKDIR /app
COPY . .
RUN go mod tidy & go build -o user-service cmd/user-service/main.go

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/user-service /
CMD ["/user-service"]