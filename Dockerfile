# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

ENV PORT_EXPOSE="8080"

WORKDIR  /app

COPY go.mod ./

RUN go mod tidy

COPY . ./

RUN go build -o /pubgo

EXPOSE $PORT_EXPOSE

CMD ["/pubgo"]