FROM golang:1.16.3-alpine

# アップデートとgitのインストール
RUN apk add --update &&  apk add git
# reflex(ホットリロード)導入
RUN go get github.com/cespare/reflex
# apiディレクトリの作成
RUN mkdir /go/src/api

# パッケージのインストール
RUN go get github.com/lib/pq

# install go tools（自動補完等に必要なツールをコンテナにインストール）
RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs \
  github.com/ramya-rao-a/go-outline \
  github.com/nsf/gocode \
  github.com/acroca/go-symbols \
  github.com/fatih/gomodifytags \
  github.com/josharian/impl \
  github.com/haya14busa/goplay/cmd/goplay \
  github.com/go-delve/delve/cmd/dlv \
  golang.org/x/lint/golint \
  golang.org/x/tools/gopls