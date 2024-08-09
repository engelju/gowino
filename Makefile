all: build run

build:
	go build -o bin/gowino

run:
	./bin/gowino

test:
	go test -v ./...