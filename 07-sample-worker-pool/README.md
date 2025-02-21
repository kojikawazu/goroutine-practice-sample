# Worker Pool Sample

## Worker Pool とは？

- **Worker Pool（ワーカープール）** は **複数の Goroutine を使って並行処理を行うパターン** です。
- 複数の **Worker** を起動し、タスクを **Channel** 経由で受け取り、効率よく処理します。
- `sync.WaitGroup` を使って **すべての Goroutine の完了を待つ** ことができます。

## Worker Pool の仕組み

1. **`jobs` チャネル** にタスクを投入
2. **複数の Worker が `jobs` からタスクを受け取り並行処理**
3. **処理結果を `results` チャネルに送信**
4. **すべての Worker が完了するまで `sync.WaitGroup` で待機**
5. **`results` から結果を取得し出力**

## Worker Pool の使い方

### 基本的な構造

```go
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
    defer wg.Done() // タスク完了時に WaitGroup から減算

    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Second) // 擬似的な処理時間
        results <- fmt.Sprintf("Job %d done by Worker %d", job, id)
    }
}
```

- `jobs <-chan int` は 受信専用チャネル
- `results chan<- string` は 送信専用チャネル
- `range jobs` でタスクを受信し、タスクがなくなるまでループ処理

## Worker Pool の注意点

### 1. チャネルを閉じる

- `jobs` チャネルはすべてのタスクを送信した後に `close(jobs)` する必要があります。
- `results` チャネルは `Worker` がすべての結果を送信した後に閉じる。

```go
close(jobs)   // タスク送信完了後
close(results) // Worker がすべての結果を送信した後
```

### 2. sync.WaitGroup の適切な管理

- `wg.Add(1)` は `Worker` を起動する前 に実行する。
- `wg.Done()` は `Worker` が終了するタイミング で実行する。

```go
wg.Add(1)
go worker(i, jobs, results, &wg)
```

