.PHONY: build run clean

build:
	go build -o store cmd/store/main.go

run:
	go run cmd/store/main.go

depend:
	go mod download && go mod verify

lint:
	golangci-lint run

clean:
	rm store