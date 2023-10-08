FROM golang:1.19-alpine AS builder

RUN apk --update upgrade && \
    apk add build-base gcc sqlite && \
    apk add --no-cache libc6-compat && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY . /app

RUN CC=gcc CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wex main.go

FROM alpine:latest

RUN apk --update upgrade && \
    apk --no-cache add ca-certificates && \
    rm -rf /var/cache/apk/*

EXPOSE 3060

COPY --from=builder /app/ /app

ENV DATABASE_HOST=postgres://postgres:secret@db_wex:5432/purchase?sslmode=disable 

WORKDIR /app

CMD ["./wex", "api"]