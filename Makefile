all: build

build:
	go build -o output/subcenter main.go

test:
	go test ./service/push

run:
	./output/subcenter
