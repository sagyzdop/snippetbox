FROM golang:alpine AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY ./cmd ./
RUN go build -o ./snippetbox ./web

FROM alpine:latest
WORKDIR /snippetbox
COPY --from=builder /build/snippetbox ./cmd/web
COPY ./ui ./ui

CMD ["./cmd/web"]
