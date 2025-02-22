# Context-Based Pipeline Sample

## Context を活用した Pipeline とは？

- Pipeline（パイプライン）は、Go 言語で **複数の処理をステージ化してデータをストリーム処理する手法** です。
- `context.Context` を利用することで、特定のタイミングで Pipeline の処理をキャンセルすることが可能になります。
- タイムアウトやキャンセル制御を組み込むことで、長時間の処理を適切に管理できます。

## Pipeline の構成

### 1. Producer（データ生成）

- `generate()` 関数が **データを作成** し、チャネルに送信。
- **キャンセル信号 (`ctx.Done()`) を監視し、途中で停止可能**。

### 2. Worker（データ処理）

- `multiplyByTwo()` 関数が **受け取ったデータを 2 倍に変換** して、次のチャネルへ送信。
- **キャンセル信号を受信すると、途中で処理を中断**。

### 3. Consumer（データ収集）

- `main()` 関数が **最終結果を受信し、出力** する。
- **データ処理がキャンセルされた場合は、途中で終了する**。

## Context を活用した Pipeline のメリット

- キャンセル可能なストリーム処理
  context.Context を使用することで、一定時間後や任意のタイミングで Pipeline を安全に停止可能。
- 長時間の処理を適切に制御
  - タイムアウト (context.WithTimeout) や手動キャンセル (context.WithCancel) に対応。
- リソースリークの防止
  - 不要な Goroutine を確実に停止し、チャネルを閉じることでメモリリークを防ぐ。

## Pipeline の注意点

### 1. ctx.Done() を適切に監視

- キャンセル可能な処理を行う場合、select 文を使って ctx.Done() を監視する。
- Goroutine のリークを防ぐため、キャンセルされたら return で処理を終了する。

```go
select {
case out <- n:
case <-ctx.Done():
    fmt.Println("Process canceled")
    return
}
```

### 2. defer close(out) でチャネルを確実に閉じる

- defer close(out) を使うことで、処理終了時に確実にチャネルを閉じる。
- これにより、Consumer 側 (range ループ) がデータを受信し終えたことを判別可能。

```go
defer close(out)
```

### 3. タイムアウトの設定

- 適切な context.WithTimeout の時間設定が重要。
- 処理時間に見合ったタイムアウト値を設定することで、不要な計算リソースを抑える。

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
```
