all: build

build:
	go build -o output/bilicenter main.go

login: build
	./output/bilicenter

run: build
	./output/bilicenter
