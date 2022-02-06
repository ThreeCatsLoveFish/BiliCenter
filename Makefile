all: build

build:
	go build -o output/sub_center main.go

test:
	go test
