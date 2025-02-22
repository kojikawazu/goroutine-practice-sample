# Timeout-Based Concurrent HTTP Requests Sample

## タイムアウト付きの並列 HTTP リクエストとは？

- このサンプルは **Go 言語で複数の HTTP リクエストを並列に実行し、タイムアウト制御を適用する方法** を示します。
- `context.WithTimeout` を利用して、**一定時間経過後にリクエストを自動キャンセル** します。
- `sync.WaitGroup` を活用して **複数の Goroutine の完了を待機** します。

## 処理の流れ

1. 複数の URL のリストを定義
2. `context.WithTimeout` で 2 秒のタイムアウトを設定
3. `fetchURL()` 関数を Goroutine として並列に実行
4. `sync.WaitGroup` を使用してすべてのリクエストが完了するまで待機
5. タイムアウトが発生した場合、リクエストをキャンセル

## メリット

- HTTP リクエストの高速化
  - 複数の URL へのリクエストを並列処理することで、レスポンスをより早く取得できる。
- タイムアウト制御
  - `context.WithTimeout` を活用することで、長時間のリクエストを制御可能。
  - サーバーのレスポンスが遅い場合に不要なリクエストをキャンセルできる。
- `sync.WaitGroup` を使った並列処理
  - Goroutine の数が増えても、すべての処理が完了するまで待機可能。

## 注意点

### 1. sync.WaitGroup の適切な管理

- `wg.Add(1)` は `go` の前に呼ぶことが推奨される。
- `defer wg.Done()` を使うことで、リクエストが完了したら WaitGroup のカウントを確実に減らせる。

```go
wg.Add(1) // `go` の前に呼ぶ
go fetchURL(ctx, url, &wg)
```

### 2. タイムアウトを適切に設定

- 処理時間に応じた適切な `context.WithTimeout` の値を設定する。

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
```

### 3. HTTP クライアントの最適化

- デフォルトの `http.DefaultClient` は Keep-Alive の設定がないため、`http.Client{Timeout:}` を活用するのが推奨される。

```go
var httpClient = &http.Client{Timeout: 5 * time.Second}

func fetchURL(ctx context.Context, url string, wg *sync.WaitGroup) {
    defer wg.Done()

    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        fmt.Printf("Failed to create request for %s: %v\n", url, err)
        return
    }

    start := time.Now()
    resp, err := httpClient.Do(req)
    if err != nil {
        fmt.Printf("Failed to fetch %s: %v\n", url, err)
        return
    }
    defer resp.Body.Close()

    fmt.Printf("Fetched %s in %v\n", url, time.Since(start))
}
```