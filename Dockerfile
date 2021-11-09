FROM golang:alpine

WORKDIR /services-api

ADD . .

RUN go mod download

ENTRYPOINT go build  && ./services-api