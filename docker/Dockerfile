FROM golang:1.23

RUN apt update && apt upgrade -y && \
    apt install -y ca-certificates

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

EXPOSE $PORT
