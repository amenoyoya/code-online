# Alpine Linux with golang image をベースにする
FROM golang:alpine

# 作業ディレクトリ: /go/ => ホスト側の ./ と共有される
WORKDIR /go/
