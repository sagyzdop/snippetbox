FROM golang:alpine AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY ./cmd ./
COPY ./internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -o ./snippetbox ./web

FROM alpine:latest
WORKDIR /snippetbox
COPY --from=builder /build/snippetbox ./cmd/web
COPY ./ui ./ui
COPY ./tls ./tls
# CMD ["./cmd/web"]
