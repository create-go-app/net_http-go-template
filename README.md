# net/http backend template for Create Go App CLI

<img align="right" width="256px" src="https://golang.org/lib/godoc/images/go-logo-blue.svg" alt="Golang logo" />

Package `net` provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets. Although the package provides access to low-level networking primitives.

Package `net/http` provides HTTP client and server implementations.

📚 [Documentation](https://golang.org/pkg/net/http/)

## Requirements

- Create Go App CLI `v0.x` ([create-go-app/cli](https://github.com/create-go-app/cli))
- Go `v1.11+` with Go Modules ([golang/download](https://golang.org/dl/))

### Optional

- Docker `19.x` ([docker/onboarding](https://hub.docker.com/?overlay=onboarding))

## Used packages

-	[joho/godotenv](https://github.com/joho/godotenv) `v1.3.0`
- [zap](https://go.uber.org/zap) `v1.13.0`
- [gorilla/mux](https://github.com/gorilla/mux) `v1.7.3`
- [json-iterator/go](https://github.com/json-iterator/go) `v1.1.9`

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

### TODO (ASAP list)

- Add more API endpoints
- Add tests
- Add jmoiron/sqlx (Postgres)
- Add markbates/pkger
- ...

## Developers

- Idea and active development by [Vic Shóstak](https://github.com/koddr) (aka Koddr).

## Project assistance

If you want to say «thank you» or/and support active development `Create Go App`:

1. Add a GitHub Star to project.
2. Twit about project [on your Twitter](https://twitter.com/intent/tweet?text=Set%20up%20a%20new%20Go%20%28Golang%29%20full%20stack%20app%20by%20running%20one%20CLI%20command%21%26url%3Dhttps%3A%2F%2Fgithub.com%2Fcreate-go-app).
3. Donate some money to project author via PayPal: [@paypal.me/koddr](https://paypal.me/koddr?locale.x=en_EN).
4. Join DigitalOcean at our [referral link](https://m.do.co/c/b41859fa9b6e) (your profit is **\$100** and we get \$25).
5. Become a sponsor.

Thanks for your support! 😘 Together, we make this project better every day.

### Sponsors

| Logo                                                                                                   | Sponsor description                                                                                                                 | URL                              |
| ------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------- | -------------------------------- |
| <img align="center" width="100px" src="https://raw.githubusercontent.com/create-go-app/cli/master/images/sponsors/1wa.co_logo.png" alt="True web artisans logo"/> | **True web artisans** — IT specialists around the world, who are ready to share their experience to solve your business objectives. | [https://1wa.co](https://1wa.co) |
|                                                                                                        | <div align="center">💡 <a href="mailto:truewebartisans@gmail.com">Want to become a sponsor too?</a></div>                           |                                  |

## License

MIT
