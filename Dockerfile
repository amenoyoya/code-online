# Based on AlpineLinux 3.12
FROM alpine:3.12

# Go言語用環境変数
ENV GOPATH /root/go
ENV GOBIN $GOROOT/bin
ENV GO111MODULE on
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN : 'install golang' && \
    apk add --update --no-cache vim git make musl-dev go curl && \
    mkdir -p ${GOPATH}/src ${GOPATH}/bin
