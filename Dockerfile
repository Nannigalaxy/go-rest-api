# Build stage
FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./main.go

# Runtime stage
FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/main .

# Copy .env file
COPY .env .

# Load environment variables from .env
RUN export $(cat .env | xargs)

EXPOSE 3000
CMD ["./main"]
