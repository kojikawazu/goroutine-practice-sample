# Parallel Batch Processing in Go

## Batch 処理

- `Go` の　`sync.WaitGroup` を活用して、複数のタスクを並列に処理するバッチ処理を実装します。
- 各バッチジョブは Goroutine を使用して並列に実行されます。
- `sync.WaitGroup` を使うことで、すべてのジョブが完了するのを待機できます。

## 処理の流れ

1. 処理する値のリスト (`values`) を定義
2. 各値に対して `processBatch()` 関数を Goroutine として並列実行
3. `sync.WaitGroup` を使用してすべてのジョブが完了するまで待機
4. すべての処理が終わったら `All batch jobs completed!` を出力

## メリット

- 並列処理によるパフォーマンス向上
  - Goroutine を使用して複数のバッチジョブを並列に処理することで、処理速度を向上できる。
- `sync.WaitGroup` を使った並列管理
  - ジョブの完了を待機することで、すべての処理が終わるまでプログラムが終了しない。
- スケール可能
  - バッチジョブの数が増えても、Goroutine を使うことで効率よく並列処理が可能。

## 注意点

### 1. `sync.WaitGroup` の適切な管理

- `wg.Add(1)` は `go` の前に呼ぶことが推奨される。
- `defer wg.Done()` を使うことで、ジョブが完了したら WaitGroup のカウントを確実に減らせる。

```go
wg.Add(1) // `go` の前に呼ぶ
go processBatch(i+1, v, &wg)
```

### 2. Worker Pool を活用した制限付き並列処理

- 大量の Goroutine を作成するとリソースが枯渇する可能性があるため、Worker Pool を使って制限するのが推奨。

```go
func workerPool(jobs <-chan int, wg *sync.WaitGroup) {
    for v := range jobs {
        processBatch(v, v, wg)
    }
}

func main() {
    values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    jobs := make(chan int, len(values))
    var wg sync.WaitGroup
    numWorkers := 3 // 3つのWorkerで並列実行

    // Workerの起動
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            workerPool(jobs, &wg)
        }()
    }

    // タスクをチャネルに送信
    for _, v := range values {
        jobs <- v
    }
    close(jobs) // 送信完了を通知

    // Workerの終了を待つ
    wg.Wait()
    fmt.Println("All batch jobs completed!")
}
```

- この実装では、3 つの Worker を使ってバッチジョブを並列処理し、リソースを効率的に管理。