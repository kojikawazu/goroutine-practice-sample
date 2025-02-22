# Cancellable Worker Pool with Context

## Cancellable Worker Pool

- Go の Worker Pool を活用し、キャンセル可能な並列処理を実装します。
- `context.WithCancel` を利用して、特定のタイミングで全 Worker を停止可能にします。
- `sync.WaitGroup` を使用して、すべての処理が完了するまで待機します。

## 処理の流れ

1. 処理するジョブの数 (`numJobs`) を定義
2. 固定数 (`numWorkers`) の Worker を起動
3. ジョブを `jobs` チャネルに送信
4. Worker は `jobs` からジョブを受信し、処理結果を `results` に送信
5. 3 秒後に `context.WithCancel` を使って全 Worker をキャンセル
6. Worker は `ctx.Done()` を受信すると処理を中断
7. Worker の終了を待機し、結果を出力

## メリット

- Worker の制御
  - ジョブが途中でキャンセルされても、すべての Worker が安全に停止可能。
- Goroutine のリーク防止
  - Worker は ctx.Done() を監視し、不要な Goroutine が実行され続けるのを防止。
- スケール可能
  - Worker の数を適切に設定することで、負荷を制御しながら並列処理が可能。

## 注意点

### 1. context.WithCancel の適切な使用

- cancel() を呼び出すことで、全 Worker にキャンセル信号が送信される。

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

### 2. Worker の適切な終了処理

- Worker は ctx.Done() を監視し、キャンセル時に処理を終了する必要がある。

```go
case <-ctx.Done():
    fmt.Printf("Worker %d: Cancelled\n", id)
    return
```

### 3. Worker Pool を活用したリソース管理

- 大量の Goroutine を作成するとリソースが枯渇する可能性があるため、Worker Pool を使って制限するのが推奨。

```go
func workerPool(ctx context.Context, jobs <-chan int, wg *sync.WaitGroup) {
    for v := range jobs {
        fmt.Printf("Processing job %d\n", v)
        time.Sleep(time.Second)
    }
}
```
