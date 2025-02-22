# Worker Pool for HTTP Requests

## 並列 HTTP リクエストの Worker Pool とは？

- Go の Worker Pool を活用し、複数の HTTP リクエストを並列処理する方法を示します。
- 複数の Worker（Goroutine）を起動し、タスクをチャネルで分配することで、効率的にリクエストを処理します。
- `sync.WaitGroup` を利用し、すべてのリクエストが完了するまで待機します。

## 処理の流れ

1. 複数の URL のリストを定義
2. 指定した数 (`numWorkers`) の Worker を起動
3. URL を `urlChan` チャネルに送信
4. Worker は `urlChan` から URL を受信し、HTTP リクエストを実行
5. 結果を `results` チャネルに送信
6. Worker がすべてのタスクを完了したら、`results` チャネルを閉じる
7. 結果を `results` から取得し、出力する

## メリット

- リクエストの並列処理
  - 複数の Worker を使って、HTTP リクエストを並列に処理することで、パフォーマンスが向上。
- Worker Pool による負荷制御
  - Worker 数 (numWorkers) を調整することで、サーバーやネットワークの負荷を抑えつつ並列処理が可能。
- Goroutine & チャネルによる効率的なタスク管理
  - urlChan を使ってリクエストを Worker に分配し、結果を results に保存することで、スレッドセーフなデータ管理が可能。

## 注意点

### 1. Worker 数 (numWorkers) の適切な設定

- Worker 数を適切に設定することで、リソースを有効活用できる。
- CPU バウンドの処理なら runtime.NumCPU() を活用し、最適な Worker 数を決定。

```go
import "runtime"

numWorkers := runtime.NumCPU() // CPU コア数を取得
```

### 2. HTTP クライアントの最適化

- http.DefaultClient を使うと接続の再利用ができないため、http.Client{} を設定するのが推奨。

```go
var httpClient = &http.Client{Timeout: 5 * time.Second}

func fetchURL(url string) {
    resp, err := httpClient.Get(url)
}
```

### 3. タイムアウトを設定

- リクエストが遅い場合、タイムアウトを設定することでリソースを有効に使える。

```go
req, err := http.NewRequest("GET", url, nil)
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
req = req.WithContext(ctx)

resp, err := httpClient.Do(req)
```
