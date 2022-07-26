all: build

build:
	go build -o output/bilicenter main.go

login: build
	./output/bilicenter login

run: build
	./output/bilicenter
