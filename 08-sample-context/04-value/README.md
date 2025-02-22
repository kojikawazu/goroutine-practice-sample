# Context WithValue Sample

## Context WithValueとは？

- `context.WithValue` は Go 言語の `context` パッケージを利用し、Goroutine にデータを渡すための仕組み です。
- `context.WithValue` を使うことで、**親 Goroutine から子 Goroutine へ値を安全に引き継ぐことができます**。
- 一般的に **リクエスト情報やユーザー ID などをコンテキスト経由で受け渡す** のに便利です。

## Context WithValue の使い方

### 基本的な書き方

```go
ctx := context.WithValue(context.Background(), "key", "value")
value := ctx.Value("key") // 値を取得
```

- `context.WithValue` を使って コンテキストにデータを格納
- `ctx.Value("key")` を使って データを取得

## Context WithValue の注意点

### 1. context.WithValue は context.WithCancel や context.WithTimeout と併用可能

- `context.WithValue` は キャンセルやタイムアウトの機能を持たない ため、`context.WithCancel` や `context.WithTimeout` と組み合わせて使うのが一般的です。

```go
ctx, cancel := context.WithCancel(context.Background())
ctx = context.WithValue(ctx, "userID", 12345)

go worker(ctx)

cancel() // Goroutine を停止
```

### 2. context.WithValue のキーには string ではなくカスタム型を推奨

- `string` をキーとして使うと、他のパッケージとの競合が起こる可能性があります。
- キーにはカスタム型を使うのが推奨 されます。

```go
type userKey string

const key userKey = "userID"

ctx := context.WithValue(context.Background(), key, 12345)
userID := ctx.Value(key)
fmt.Println("User ID:", userID)
```

### 3. context.WithValue は大量のデータを格納する用途には向かない

- `context.WithValue` は 軽量なデータ（リクエスト ID やユーザー情報） の受け渡しに適しています。
- 大きなデータ（JSON, 配列, DB の結果など）は推奨されません。

```go
ctx := context.WithValue(context.Background(), "userData", hugeData)
```

- 代替手段: 必要なデータのみ渡し、詳細情報は データストア (DB, キャッシュ) で管理 する。
