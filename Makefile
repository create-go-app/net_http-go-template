.PHONY: clean test security build run swag

APP_NAME = apiserver
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = postgres://postgres:password@localhost/postgres?sslmode=disable

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

linter:
	golangci-lint run

test: security
	go test -cover ./...

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

docker.build:
	docker build -t net_http-go-template .

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.run: docker.network docker.postgres swag docker.net_http migrate.up

docker.stop:
	docker stop dev-net_http dev-postgres

docker.net_http: docker.build
	docker run --rm -d \
		--name dev-net_http \
		--network dev-network \
		-p 5000:5000 \
		net_http-go-template

docker.postgres:
	docker run --rm -d \
		--name dev-postgres \
		--network dev-network \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=postgres \
		-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres

swag:
	swag init