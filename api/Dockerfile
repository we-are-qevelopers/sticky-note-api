FROM golang:1.16.3-alpine

# アップデートとgitのインストール
RUN apk add --update &&  apk add git
# reflex(ホットリロード)導入
RUN go get github.com/cespare/reflex
# apiディレクトリの作成
RUN mkdir /go/src/api

