FROM golang:latest

LABEL maintainer="mudrinic.mare@gmail.com"

ADD . /go/src/github.com/xmudrii/etcdstorage-apiserver

WORKDIR /go/src/github.com/xmudrii/etcdstorage-apiserver

RUN make compile

ENTRYPOINT ["./bin/etcdstorage-apiserver"]