# Context WithDeadline Sample

## Context WithDeadlineとは？

- `context.WithDeadline` は、指定した日時 (`time.Time`) になると自動的にキャンセルされる `context` を作成する仕組み です。
- `context.WithTimeout` と似ていますが、絶対時間 (`time.Time`) を指定する点が異なります。
- API リクエスト、バッチ処理、一定時間後に終了させたいタスクなどに利用されます。

## Context WithDeadline の使い方

### 基本的な書き方

```go
deadline := time.Now().Add(2 * time.Second) // 2秒後の時間を指定
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel() // メモリリーク防止

select {
case <-ctx.Done():
    fmt.Println("Deadline reached, stopping process")
}
```

- `context.WithDeadline(context.Background(), deadline)`
  - 指定した deadline に達すると、自動でキャンセルされる context を作成
- `defer cancel()`
  - 必ず defer cancel() を呼んでリソースを解放
- `ctx.Done()`
  - 締め切り (deadline) が過ぎると <-ctx.Done() に通知される

## Context WithDeadline の注意点

### 1. defer cancel() を必ず呼ぶ

- `context.WithDeadline` を使う場合、`defer cancel()` を必ず呼ぶことで、リソースを適切に解放する。

```go
ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
defer cancel() // メモリリーク防止
```

### 2. タイムアウトを select で監視

- `ctx.Done()` を適切に監視しないと、Goroutine は締め切り (deadline) 以降も動作を続けてしまう。

```go
select {
case <-ctx.Done(): // 締め切り時間が到達したらキャンセル
    fmt.Println("Deadline reached, stopping process")
}
```

### 3. context.WithTimeout との違い

- `context.WithDeadline` は 絶対時間 (time.Time) を指定
- `context.WithTimeout` は 相対時間 (time.Duration) を指定

```go
// WithDeadline
deadline := time.Now().Add(2 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)

// WithTimeout
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
```
