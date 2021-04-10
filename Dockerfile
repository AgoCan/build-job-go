# stage 1: build src code to binary
FROM golang:1.16.3-alpine3.13 as builder

ENV GOPROXY=https://goproxy.cn
ENV GOPATH=$HOME/go
ENV GOBIN=$HOME/go/bin
ENV PATH=$PATH:$GOPATH/bin

COPY . /app/

RUN cd /app && go build -o build-job .

# stage 2: use alpine as base image
FROM alpine:3.13

RUN apk update && \
    apk --no-cache add tzdata ca-certificates && \
    cp -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*
# 使用--from的参数做到拷贝使用
COPY --from=builder /app/build-job /build-job
COPY config/config.yaml /etc/
CMD ["/build-job -c /etc/config.yaml"]