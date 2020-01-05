# net_http-go-template

Backend template with built-in `net/http` (pkg/net/http).

## Project structure

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
