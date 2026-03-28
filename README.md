# go-graphql-test
## 概要
GoでGraphQLのサーバーを構築するためのサンプルプロジェクトです。

## 使い方
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
go run main.go
```

## 使用ライブラリ
- [graphql-go](https://github.com/99designs/gqlgen): GraphQLサーバーの実装に使用しています。
- [ent](https://entgo.io/ent/cmd/ent): データベースのORMとして使用しています。