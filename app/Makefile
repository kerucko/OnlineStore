.PHONY: build run clean

build:
	go build -o store cmd/store/main.go

run:
	go run cmd/store/main.go

depend:
	go mod download && go mod verify

lint:
	golangci-lint run

migrate:
	goose -dir ./migrations postgres "user=postgres dbname=postgres password=postgres host=db" up

down:
	goose -dir ./migrations postgres "user=postgres dbname=postgres password=postgres host=db" down

reset:
	goose -dir ./migrations postgres "user=postgres dbname=postgres password=postgres host=db" reset

clean: reset
	rm store
