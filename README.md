# net/http backend template for [Create Go App CLI](https://github.com/create-go-app/cli)

<img src="https://img.shields.io/badge/Go-1.17+-00ADD8?style=for-the-badge&logo=go" alt="go version" />&nbsp;<a href="https://goreportcard.com/report/github.com/create-go-app/net_http-go-template" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>&nbsp;<img src="https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none" alt="license" />

Package `net` provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets. Although the package provides access to low-level networking primitives.

Package [`net/http`](https://golang.org/pkg/net/http/) provides HTTP client and server implementations.

## ⚡️ Quick start

1. Create a new project with Fiber:

```bash
cgapp create

# Choose a backend framework:
# > net/http
#   fiber
#   chi
```

2. Rename `.env.example` to `.env` and fill it with your environment values.
3. Install [Docker](https://www.docker.com/get-started) and the following useful Go tools to your system:

   - [golang-migrate/migrate](https://github.com/golang-migrate/migrate#cli-usage) for apply migrations
   - [github.com/swaggo/swag](https://github.com/swaggo/swag) for auto-generating Swagger API docs
   - [github.com/securego/gosec](https://github.com/securego/gosec) for checking Go security issues
   - [github.com/go-critic/go-critic](https://github.com/go-critic/go-critic) for checking Go the best practice issues
   - [github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint) for checking Go linter issues

4. Run project by this command:

```bash
make docker.run
```

5. Go to API Docs page (Swagger): [127.0.0.1:5000/swagger/index.html](http://127.0.0.1:5000/swagger/index.html)

![Screenshot](https://user-images.githubusercontent.com/11155743/112716623-76bc0800-8ef8-11eb-80db-48edcbdacf36.png)

## 📦 Used packages

| Name                                                                  | Version   | Type       |
| --------------------------------------------------------------------- | --------- | ---------- |
| [net/http](https://golang.org/pkg/net/http/)                          | `latest`  | core       |
| [gorilla/mux](https://github.com/gorilla/mux)                         | `v1.8.0`  | router     |
| [auth0/go-jwt-middleware](https://github.com/auth0/go-jwt-middleware) | `v1.0.1`  | middleware |
| [swaggo/http-swagger](https://github.com/swaggo/http-swagger)         | `v1.1.2`  | middleware |
| [stretchr/testify](https://github.com/stretchr/testify)               | `v1.7.0`  | tests      |
| [golang-jwt/jwt](https://github.com/golang-jwt/jwt)                   | `v4.2.0`  | auth       |
| [joho/godotenv](https://github.com/joho/godotenv)                     | `v1.4.0`  | config     |
| [jmoiron/sqlx](https://github.com/jmoiron/sqlx)                       | `v1.3.4`  | database   |
| [jackc/pgx](https://github.com/jackc/pgx)                             | `v4.14.1` | database   |
| [swaggo/swag](https://github.com/swaggo/swag)                         | `v1.7.6`  | utils      |
| [google/uuid](https://github.com/google/uuid)                         | `v1.3.0`  | utils      |
| [go-playground/validator](https://github.com/go-playground/validator) | `v10.9.0` | utils      |

## 🗄 Template structure

### ./app

**Folder with business logic only**. This directory doesn't care about _what database driver you're using_ or _which caching solution your choose_ or any third-party things.

- `./app/controllers` folder for functional controllers (used in routes)
- `./app/models` folder for describe business models and methods of your project
- `./app/queries` folder for describe queries for models of your project
- `./app/validators` folder for describe validators for models fields

### ./docs

**Folder with API Documentation**. This directory contains config files for auto-generated API Docs by Swagger.

### ./pkg

**Folder with project-specific functionality**. This directory contains all the project-specific code tailored only for your business use case, like _configs_, _middleware_, _routes_ or _utils_.

- `./pkg/configs` folder for configuration functions
- `./pkg/middleware` folder for add middleware (Fiber built-in and yours)
- `./pkg/routes` folder for describe routes of your project
- `./pkg/utils` folder with utility functions (server starter, error checker, etc)

### ./platform

**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual project, like _setting up the database_ or _cache server instance_ and _storing migrations_.

- `./platform/database` folder with database setup functions (by default, PostgreSQL)
- `./platform/migrations` folder with migration files (used with [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool)

## ⚙️ Configuration

```ini
# .env

# Server settings:
SERVER_HOST="0.0.0.0"
SERVER_PORT=5000
SERVER_READ_TIMEOUT=60

# JWT settings:
JWT_SECRET_KEY="secret"
JWT_REFRESH_KEY="refresh"

# Database settings:
DB_HOST="cgapp-postgres"
DB_PORT=5432
DB_USER="postgres"
DB_PASSWORD="password"
DB_NAME="postgres"
DB_SSL_MODE="disable"
DB_MAX_CONNECTIONS=100
DB_MAX_IDLE_CONNECTIONS=10
DB_MAX_LIFETIME_CONNECTIONS=2
```

## ⚠️ License

Apache 2.0 &copy; [Vic Shóstak](https://shostak.dev/) & [True web artisans](https://1wa.co/).
