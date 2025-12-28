.DEFAULT_GOAL := build

build:
		@rm -rf ./build
		@mkdir ./build
		go build -o ./build/tcp-port-checker ./cmd/MakSec-TCP-Port-Check/main.go

run: build
		./build/tcp-port-checker