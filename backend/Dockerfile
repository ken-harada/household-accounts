FROM golang:1.22-alpine AS builder

WORKDIR /var/www/backend

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]
