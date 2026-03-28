# go-graphql-test
## 概要
GoでGraphQLのサーバーを構築するためのサンプルプロジェクトです。

## 使い方
### Dockerを使用する場合

1. Dockerイメージをビルドします。
```bash
docker build -t go-graphql-test .
```

2. コンテナを起動します。
```bash
docker run -p 8080:8080 go-graphql-test
```

3. ブラウザで `http://localhost:8080` にアクセスして、GraphQL Playgroundを使用してクエリを実行できます。

### Go環境が既にセットアップされている場合
1. リポジトリをクローンします。

```bash
git clone https://github.com/yourusername/go-graphql-test.git
```

2. プロジェクトディレクトリに移動します。

```bash
cd go-graphql-test
```

3. 依存関係をインストールします。

```bash
go mod tidy
```

4. サーバーを起動します。

```bash
go run server.go
```

5. ブラウザで `http://localhost:8080` にアクセスして、GraphQL Playgroundを使用してクエリを実行できます。

## 使用ライブラリ
- [graphql-go](https://github.com/99designs/gqlgen): GraphQLサーバーの実装に使用しています。
- [ent](https://entgo.io/ent/cmd/ent): データベースのORMとして使用しています。