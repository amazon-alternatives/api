# Build stage
FROM golang:1.15 AS builder

WORKDIR /api

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main ./cmd/api/main.go

# Final stage
FROM alpine:3

COPY --from=builder /main ./
RUN chmod +x ./main

ENTRYPOINT ["./main"]
