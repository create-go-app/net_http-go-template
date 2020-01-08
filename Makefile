.PHONY: run

run:
	go run ./cmd/apiserver/*.go

build:
	rm -rf ./app \
	&& CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
	go build -o ./app/apiserver ./cmd/apiserver/*.go
	@echo "[✔️] Backend was builded!"

.DEFAULT_GOAL := run
