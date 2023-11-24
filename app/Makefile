.PHONY: build run clean

build:
	go build cmd/store/main.go

run:
	go run cmd/store/main.go

lint:
	golangci-lint run

clean:
	del main.exe # винда сосет

depend:
	go mod download && go mod verify