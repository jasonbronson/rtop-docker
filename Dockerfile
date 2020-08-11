
FROM golang:1.14-alpine

RUN apk update && apk add git bash build-base curl
WORKDIR /go/src
RUN mkdir -p rtop
WORKDIR /go/src/rtop

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

COPY ./ .