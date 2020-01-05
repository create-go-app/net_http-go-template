.PHONY: run

run:
	go run ./*.go

build:
	go build -v ./main.go

.DEFAULT_GOAL := run
