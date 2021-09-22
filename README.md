# Go言語｜入門

## Environment

- OS:
    - Windows 10
    - Ubuntu 18.04
- Editor:
    - VS Code

***

## Docker

Goをインストールして使っても良いが、Dockerを使えばインストール不要で使える

### Structure
```bash
./ # カレントディレクトリ = 作業ディレクトリ（appコンテナの /go/ ディレクトリにマウントされる）
|_ Dockerfile # appコンテナビルド設定
|_ docker-compose.yml # appコンテナ: golang:alpine
```

### Usage
基本的に Dockerコンテナを以下のようにコマンドとして使うことを想定している

```bash
# 利用時にコンテナ起動＆コマンド実行 => コマンド完了後コンテナ削除（--rm オプション）
## ※初回起動時のみコンテナイメージのダウンロード＆ビルドに時間がかかる

# 例: test/main.go をスクリプトとして実行
## go run test/main.go
$ docker-compose run --rm go run test/main.go

# 例: test/main.go を main 実行ファイルにコンパイル
## go build -o test/main test/main.go
$ docker-compose run --rm go build -o test/main test/main.go
```

***

## Installation

コンピュータにインストールして使う場合（for Windows10, Ubuntu 18.04）

### on Windows 10
[Chocolatey](https://chocolatey.org/) を使うのが楽

`Win + X` |> `A` キー => 管理者権限PowerShell起動

```powershell
# install golang
> choco install -y golang

## => installed to C:\Program Files\Go

# -- Restart PowerShell for reflecting the environmental variables

# confirm version
> go version
go version go1.17.1 windows/amd64
```

### on Ubuntu 18.04
[最新の公式パッケージ](https://golang.org/dl/) をダウンロードしてインストールする

```bash
# install in home directory
$ cd ~

# download golang binary
$ wget -O - https://dl.google.com/go/go1.17.1.linux-amd64.tar.gz | tar zxvf -

# create symbolic link to /usr/local/bin/ => enable `go`, `gofmt` command
$ sudo ln -s ~/go/bin/go /usr/local/bin/go
$ sudo ln -s ~/go/bin/gofmt /usr/local/bin/gofmt

# confirm version
$ go version
go version go1.17.1 linux/amd64
```

***

## 入門

[A Tour of Go](https://go-tour-jp.appspot.com/list) を参考に

### Go言語のコンパイル
コンパイルは基本的に `go build <ソースファイル>` で可能

```bash
# 入門プロジェクト: tuto/ ディレクトリ
$ cd tuto

# 01_hello.go をコンパイル
$ go build 01_hello.go

## => ソースコードに問題なければ 01_hello (01_hello.exe) というバイナリファイルが生成される

# 実行
$ ./01_hello

## => "Hello, 世界"
```
