FROM golang:1.19 AS builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o wex-coding-challenge main.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN apk --update upgrade && \
    apk add tzdata && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=builder /app/ ./

EXPOSE 3060

CMD ["./wex-coding-challenge", "api"]
