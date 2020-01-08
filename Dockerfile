FROM golang:alpine AS builder

WORKDIR /app
COPY . .
VOLUME /configs
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go mod download \
	&& go build -o ./apiserver ./cmd/apiserver/*.go

FROM scratch

COPY --from=builder /app/apiserver /app/
ENTRYPOINT /app/apiserver
