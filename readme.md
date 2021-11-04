### 起動方法
1. プロジェクトのルートで`docker-compose up`

#### vscodeのRemote Containersを使って開発する方法
[wip]
1. vscodeワークスペースの右の、リモートエクスプローラーを選択
2. リモートエクスプローラーのサイドバー上部のセレクターからcontainersを選択
3. container/dev Containersからコンテナの中に入って開発する
4. その際に、/go/src/apiディレクトリを開く

コンテナ入ってから、goのserverのログを標準出力させるコマンド
(apiサーバーdockerコンテナのプロセス１が、ginサーバーのプロセス)
```
cat /proc/1/fd/1 
```