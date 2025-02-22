# Parallel HTTP Requests Sample

## 並列 HTTP リクエストとは？

- 複数の HTTP リクエストを並列に処理することで、パフォーマンスを向上させる手法です。
- `Go` の `sync.WaitGroup` を活用することで、すべてのリクエストが完了するまで待機できます。
- `Goroutine` を利用することで、非同期にリクエストを実行できます。

## 並列 HTTP リクエストの仕組み

1. リクエストのリストを用意

- `urls := []string{"URL1", "URL2", "URL3"}` のように対象の URL をリストで管理。

2. `fetchURL()` 関数を Goroutine で実行

- `http.Get(url)` を使ってリクエストを実行。
- `sync.WaitGroup` を使って、リクエストが完了するのを待つ。

3. `wg.Wait()` で全リクエストの完了を待機

- Goroutine 内で `wg.Done()` を呼び、完了を通知。

## メリット

- HTTP リクエストの高速化
  - 複数の URL へのリクエストを並列処理することで、レスポンスをより早く取得できる。
- sync.WaitGroup を使った並列処理
  - Goroutine の数が増えても、すべての処理が完了するまで待機可能。
- シンプルなコードで高効率
  - シンプルな Go の Goroutine と sync.WaitGroup の組み合わせで並列処理を実装できる。

## 注意点

### 1. sync.WaitGroup の適切な管理

- `wg.Add(1)` は `go` の前に呼ぶことが推奨される。
- `defer wg.Done()` を使うことで、リクエストが完了したら WaitGroup のカウントを確実に減らせる。

```go
wg.Add(1) // `go` の前に呼ぶ
go fetchURL(url, &wg)
```

## 2. タイムアウトを設定

- ネットワークの問題でリクエストが遅延する可能性があるため、タイムアウトを設定する。

```go
client := http.Client{
    Timeout: 3 * time.Second, // 3秒でタイムアウト
}

resp, err := client.Get(url)
```

## 3. HTTP クライアントの再利用

- デフォルトの `http.Get()` は毎回新しい HTTP クライアントを作成するため、パフォーマンスが低下する可能性がある。
- 以下のように `http.Client` を使って接続を再利用すると、効率が向上する。

```go
var httpClient = &http.Client{Timeout: 5 * time.Second}

func fetchURL(url string, wg *sync.WaitGroup) {
    defer wg.Done()

    start := time.Now()
    resp, err := httpClient.Get(url)
    if err != nil {
        fmt.Printf("Failed to fetch %s: %v\n", url, err)
        return
    }
    defer resp.Body.Close()

    fmt.Printf("Fetched %s in %v\n", url, time.Since(start))
}
```
