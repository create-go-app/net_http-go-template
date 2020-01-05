# net/http backend template for Create Go App CLI

<img align="right" width="256px" src="https://golang.org/lib/godoc/images/go-logo-blue.svg" alt="Golang logo" />

Package `net` provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets. Although the package provides access to low-level networking primitives.

Package `net/http` provides HTTP client and server implementations.

📚 [Documentation](https://golang.org/pkg/net/http/)

## Requirements

- Create Go App CLI `0.x` ([create-go-app/cli](https://github.com/create-go-app/cli))
- Go `1.11+` with Go Modules ([golang/download](https://golang.org/dl/))

### Optional

- Docker `19.x` ([docker/onboarding](https://hub.docker.com/?overlay=onboarding))

## Template structure

```console
foo@bar:backend$ tree .
.
├── .dockerignore
├── .editorconfig
├── .env.example
├── .gitignore
├── Dockerfile
├── Makefile
├── README.md
├── LICENSE
├── go.mod
├── go.sum
├── cmd
│   └── apiserver
│       └── main.go
└── internal
    └── apiserver
        ├── apiserver.go
        ├── checker.go
        ├── config.go
        ├── logger.go
        ├── middleware.go
        └── routes.go

4 directories, 18 files
```
