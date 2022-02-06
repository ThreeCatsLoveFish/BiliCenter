all: build

build:
	go build -o output/subcenter main.go

test:
	go test
