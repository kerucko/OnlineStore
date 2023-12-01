FROM golang:1.21.3-alpine

WORKDIR /usr/src/app

RUN apk --no-cache add bash make
RUN apk --update add postgresql-client

COPY go.mod go.sum ./
RUN go mod download \
    && go mod verify \
    && go install github.com/pressly/goose/v3/cmd/goose@latest

COPY ./ ./
RUN make build