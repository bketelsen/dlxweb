FROM golang:1.17.2-alpine3.14 as builder

ENV GOBIN /go/bin

WORKDIR /dlxweb

COPY . .

RUN apk add --update --no-cache \
    make \
    gcc \
    libc-dev \
    npm

RUN make deps \
     tailwind \
     && make build 

FROM alpine:3.14

COPY --from=builder /dlxweb/bin/dlx-linux-amd64 /usr/bin/dlxweb

