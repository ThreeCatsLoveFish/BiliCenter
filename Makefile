all: build

build:
	go build -o output/subcenter main.go

run: build
	./output/subcenter
