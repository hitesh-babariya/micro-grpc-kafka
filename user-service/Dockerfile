FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY user-service ./user-service
COPY common ./common
COPY proto ./proto
COPY events ./events

#COPY . .

WORKDIR /app/user-service
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/user-service/app .
EXPOSE 50051
CMD ["./app"]
