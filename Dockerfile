FROM golang:1.22.1-alpine3.19 as server-build

WORKDIR /go/src/todo-list-server

COPY . .

RUN apk upgrade --update && \
    apk --no-cache add git