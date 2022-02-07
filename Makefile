all: build

build:
	go build -o output/subcenter main.go

test:
	go test ./push

run:
	./output/subcenter
