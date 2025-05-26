FROM golang:alpine AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY --chown=app:app ./ui ./ui
RUN CGO_ENABLED=0 GOOS=linux go build -o ./snippetbox ./web

FROM alpine:latest
WORKDIR /snippetbox
COPY --from=builder /build/snippetbox ./cmd/web
COPY ./tls ./tls
# CMD ["./cmd/web"]
