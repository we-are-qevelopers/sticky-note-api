FROM postgres:11-alpine
ENV LANG ja_JP.utf8
# 初期化用 sql ファイルを、コンテナ内の所定のディレクトリにコピー
COPY ./postgres/init/*.sql /docker-entrypoint-initdb.d/