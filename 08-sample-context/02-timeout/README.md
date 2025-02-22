# Context WithTimeout Sample

## Context WithTimeoutとは？

- `context.WithTimeout` は、指定した時間が経過すると自動的にキャンセルされる `context` を作成する仕組みです。
- タイムアウト時間が経過すると、関連する **すべての Goroutine にキャンセル信号 (`ctx.Done()`) が送信** されます。
- API リクエスト、DB クエリ、長時間の処理などにおいて、一定時間内に完了しない場合に強制停止させるために使用されます。

---

## Context WithTimeout の使い方

### 基本的な書き方

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel() // メモリリーク防止

select {
case <-ctx.Done():
    fmt.Println("Timeout reached, stopping process")
}
```

- `context.WithTimeout(context.Background(), 2*time.Second)`
  - 2秒後に自動でキャンセルされる context を作成
- `defer cancel()`
  - 必ず defer cancel() を呼んでリソースを解放
- `ctx.Done()`
  - タイムアウトが発生すると <-ctx.Done() に通知される

## Context WithTimeout の注意点

### 1. defer cancel() を必ず呼ぶ

- `context.WithTimeout` を使う場合、`defer cancel()` を必ず呼ぶことで、リソースを適切に解放する。

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel() // メモリリーク防止
```

### 2. タイムアウトを select で監視

- `ctx.Done()` を適切に監視しないと、Goroutine はタイムアウト後も動作を続けてしまう。

```go
select {
case <-ctx.Done(): // タイムアウト時にキャンセル
    fmt.Println("Timeout reached, stopping process")
}
```

### 3. context.WithCancel と組み合わせる

- `context.WithCancel` と組み合わせることで、より柔軟な制御が可能

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
ctx, cancelManual := context.WithCancel(ctx) // 手動キャンセルも追加
defer cancel()
defer cancelManual()
```
