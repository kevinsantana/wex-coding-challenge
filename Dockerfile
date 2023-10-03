FROM golang:1.19 AS builder

RUN mkdir -p $GOPATH/src/github.com/kevinsantana/purchase

COPY go.mod $GOPATH/src/github.com/kevinsantana/purchase/
COPY go.sum $GOPATH/src/github.com/kevinsantana/purchase/
COPY . $GOPATH/src/github.com/kevinsantana/purchase/

RUN go build -o $GOPATH/bin/purchase $GOPATH/src/github.com/kevinsantana/purchase/


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN apk --update upgrade && \
    apk add tzdata && \
    rm -rf /var/cache/apk/*

EXPOSE 3060

COPY --from=builder /go/src/github.com/kevinsantana/purchase/ /bin/purchase/

COPY --from=builder /go/bin/purchase /bin/

WORKDIR /bin

CMD ["purchase", "run"]
