# Pipeline Sample

## Pipeline とは？

- Pipeline（パイプライン）は、Go 言語でデータを **複数のステージ（処理段階）** に渡して処理する手法です。
- 各ステージは **チャネルを介してデータを受け取り、次のステージへ送信** します。
- データのストリーム処理や並列処理を効率化するのに適しています。

---

## Pipeline の仕組み

### 1. Producer（データ生成）
   - `generate()` 関数が **データを作成** し、チャネルに送信。

### 2. Worker（データ処理）
   - `multiplyByTwo()` 関数が **受け取ったデータを 2 倍に変換** して、次のチャネルへ送信。

### 3. Consumer（データ収集）
   - `main()` 関数が **最終結果を受信し、出力** する。

---

## Pipeline の基本的な使い方

### 1. データを生成する（Producer）

```go
func generate(numbers ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range numbers {
            out <- n
        }
        close(out)
    }()
    return out
}
```

- numbers のリストをチャネル out に送信
- すべてのデータを送信したら close(out) でチャネルを閉じる

### 2. データを処理する（Worker）

```go
func multiplyByTwo(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * 2
        }
        close(out)
    }()
    return out
}
```

- 受信したデータを 2 倍に変換し、新しいチャネル out に送信
- 処理が完了したら close(out) でチャネルを閉じる

## Pipeline の注意点

### 1. チャネルを必ず close() する

- close() をしないと、データを受け取る側 (range ループ) がブロックされ続けてしまう。
  
```go
close(out) // すべてのデータを送信したらチャネルを閉じる
```

### 2. goroutine のリークを防ぐ

- goroutine のリークを防ぐため、不要になったチャネルの送信側を適切に終了する。
- もしエラーが発生した場合に defer close(out) を使うことで安全に終了可能。

```go
go func() {
    defer close(out) // goroutine を確実に終了
    for n := range in {
        out <- n * 2
    }
}()
```