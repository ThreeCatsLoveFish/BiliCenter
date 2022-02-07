all: build

build:
	go build -o output/subcenter main.go

test:
	go test

run:
	./output/subcenter
