# Context WithCancel Sample

## Context WithCancelとは？

- `context.WithCancel` は **Go 言語で Goroutine を安全に停止するための仕組み** です。
- 通常の Goroutine は `main` 関数が終了しても動作を続けてしまう可能性があります。
- `context.WithCancel` を使うことで、**明示的に Goroutine を終了させることができます**。

## **Context WithCancelの使い方**

### **基本的な書き方**

```go
ctx, cancel := context.WithCancel(context.Background())

go func() {
    for {
        select {
        case <-ctx.Done(): // キャンセル信号を受信したら終了
            fmt.Println("Goroutine stopped")
            return
        default:
            fmt.Println("Working...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}()

time.Sleep(2 * time.Second)
cancel() // すべての Goroutine を停止
```

`ctx.Done()` を監視することで、Goroutine を安全に終了できます。

## Context WithCancelの注意点

### 1. ctx.Done() を必ずチェックする

`Goroutine` 内で `ctx.Done()` をチェックしないと、キャンセル信号を受信できずに 無限ループする 可能性があります。

```go
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done(): // キャンセル信号を受信
            fmt.Println("Worker stopped")
            return
        default:
            fmt.Println("Worker is running...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}
```

### 2. cancel() を必ず呼ぶ

`context.WithCancel` を作成した場合は、`cancel()` を必ず呼び出す 必要があります。
`cancel()` を呼ばないと、メモリリークの原因になることがあります。

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // メモリリーク防止
```

