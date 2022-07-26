all: run

BIN=./output/bilicenter

build:
	go build -o $(BIN) main.go

login: build
	$(BIN) login

run: build
	$(BIN)
