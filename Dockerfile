FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o apiserver cmd/apiserver/*.go

FROM scratch

COPY --from=builder ["/app/apiserver", "/app/configs/apiserver.yml", "/app/"]
ENTRYPOINT ["/app/apiserver", "-config-path", "/app/apiserver.yml"]
