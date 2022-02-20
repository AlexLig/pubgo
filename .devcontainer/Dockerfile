# syntax=docker/dockerfile:1

FROM golang:1.17

ENV SERVER_PORT=8080

WORKDIR /app

RUN go install github.com/go-delve/delve/cmd/dlv@latest

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./

RUN go mod download && go mod verify

COPY . .

EXPOSE $SERVER_PORT
