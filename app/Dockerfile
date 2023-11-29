FROM golang:1.21.3-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./ ./
RUN go build -o store cmd/store/main.go

CMD ["./store"]