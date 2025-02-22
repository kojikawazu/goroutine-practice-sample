# Worker Pool for Batch Processing in Go

## Batch Worker Pool

- `Go` の `Worker Pool` を活用して、複数のジョブを並列に処理する方法を示します。
- 固定数の `Worker`（`Goroutine`）を起動し、タスクをチャネルで分配します。
- `sync.WaitGroup` を利用し、すべてのジョブが完了するまで待機します。

## 処理の流れ

1. 処理するジョブの数 (`numJobs`) を定義
2. 固定数 (`numWorkers`) の Worker を起動
3. ジョブを `jobs` チャネルに送信
4. Worker は `jobs` からジョブを受信し、処理結果を `results` に送信
5. Worker の処理が完了すると `results` チャネルに結果を送信
6. `results` チャネルをクローズし、処理結果を出力

## メリット

- リソースの効率的な活用
  - Worker 数を固定することで、システムリソースの過負荷を防ぎつつ並列処理が可能。
- Goroutine & チャネルによる効率的なタスク管理
  - `jobs` を使ってリクエストを Worker に分配し、結果を `results` に保存することで、スレッドセーフなデータ管理が可能。
- ジョブが大量でもスケール可能
  - ジョブの数が増えても、Worker を増やすことで処理時間を短縮できる。

## 注意点

### 1. Worker 数 (numWorkers) の適切な設定

- Worker 数が少なすぎると並列処理の効果が低下し、多すぎるとリソースの消費が増加する。
- CPU バウンドの処理なら `runtime.NumCPU()` を活用し、最適な Worker 数を決定。

```go
import "runtime"

numWorkers := runtime.NumCPU() // CPU コア数を取得
```

### 2. sync.WaitGroup の適切な管理

- `wg.Add(1)` は `go` の前に呼ぶことが推奨される。
- `defer wg.Done()` を使うことで、ジョブが完了したら WaitGroup のカウントを確実に減らせる。

```go
wg.Add(1) // `go` の前に呼ぶ
go processBatch(i+1, v, &wg)
```

### 3. Worker Pool を活用した制限付き並列処理

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

- この実装では、3 つの Worker を使ってジョブを並列処理し、リソースを効率的に管理。