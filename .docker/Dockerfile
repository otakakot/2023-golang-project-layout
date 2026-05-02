FROM golang:1.21.4-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

ARG workdir

WORKDIR /app/${workdir}

CMD ["air"]
