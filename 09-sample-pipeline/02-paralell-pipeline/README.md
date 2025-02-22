# Parallel Pipeline Sample

## Parallel Pipeline（並列パイプライン）とは？

- Parallel Pipeline（並列パイプライン）は、**複数の Goroutine（Worker）を利用してデータ処理を並列化する手法** です。
- 各 Worker が **チャネルを介してデータを受信し、並行して処理を行い、結果を出力チャネルへ送信** します。
- **負荷の高い処理を複数の Worker に分散させることで、処理時間を短縮** できます。

## Pipeline の構成

### 1. Producer（データ生成）
   - `generate()` 関数が **データを作成** し、チャネルに送信。

### 2. Worker Pool（並列処理）
   - `multiplyByTwo()` 関数が **複数の Worker を起動し、受け取ったデータを 2 倍に変換** して、次のチャネルへ送信。

### 3. Consumer（データ収集）
   - `main()` 関数が **最終結果を受信し、出力** する。

## Parallel Pipeline のメリット

- 並列処理によるパフォーマンス向上
  複数の Worker を起動し、データを並列に処理することで 処理速度が向上 する。
- Goroutine を活用した効率的な並列処理
  Worker ごとに Goroutine を起動し、データをバランスよく処理できる。
- チャネルを活用したスレッドセーフなデータ共有
  チャネル (chan) を利用することで、競合なしでデータをやり取り できる。

## Parallel Pipeline の注意点

### 1. チャネルを必ず close() する

- close() をしないと、データを受け取る側 (range ループ) がブロックされ続けてしまう。

```go
close(out) // すべてのデータを送信したらチャネルを閉じる
```

### 2. sync.WaitGroup を適切に管理

- Worker が すべての処理を完了してから出力チャネルを閉じる ことが重要。

```go
go func() {
    wg.Wait()
    close(out) // すべての Worker が終了してからチャネルを閉じる
}()
```

### 3. Worker 数の最適化

- Worker 数を適切に調整することで、CPU を最大限活用できる。
- CPU バウンドな処理（計算が重い処理） の場合、runtime.NumCPU() を使って最適な Worker 数を決定 する。

```go
numWorkers := runtime.NumCPU() // CPU コア数を取得
```