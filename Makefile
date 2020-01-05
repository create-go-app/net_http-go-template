.PHONY: run

run:
	go run ./cmd/apiserver/*.go

build:
	rm -rf ./app \
	&& go build -o ./app/apiserver ./cmd/apiserver/*.go
	@echo "[✔️] Backend was builded!"

.DEFAULT_GOAL := run
