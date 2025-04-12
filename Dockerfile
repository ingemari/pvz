FROM golang:1.23 AS builder

WORKDIR ${GOPATH}/pvz/
COPY . ${GOPATH}/pvz/

RUN go build -o /build ./cmd \
    && go clean -cache -modcache

EXPOSE 8080

CMD ["/build"]