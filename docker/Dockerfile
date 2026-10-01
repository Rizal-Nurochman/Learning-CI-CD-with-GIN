FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" \
    -o /app/server ./cmd

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8000
CMD ["./server"]