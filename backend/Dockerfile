# Build the Go application
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/apiserver

# Run the application
FROM scratch

COPY --from=builder /main /main