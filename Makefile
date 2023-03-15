.PHONY: build

run:
	go run src/main.go

build:
	go build -o ./executable ./...

test:
	go test -v ./...