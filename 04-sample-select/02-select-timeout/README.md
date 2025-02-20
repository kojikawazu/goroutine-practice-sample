# Select Timeout Sample

## **Select Timeoutとは？**

- Go 言語では、通常 `select` 文を使って複数のチャネルの入力を待ちますが、  
- データが来ない場合に **タイムアウト** させる仕組みとして `time.After` を活用できます。

## **`select timeout` の仕組み**

- ある処理が **指定時間以内** に完了しなければ、タイムアウト処理を実行する。
- `time.After(duration)` を使って、**指定時間が経過すると値を送信するチャネル** を作成する。

## **Select Timeoutの使い方**

基本形は以下の通りです。

```go
select {
case msg := <-ch:
    fmt.Println("Received:", msg) // ch からデータを受信
case <-time.After(2 * time.Second):
    fmt.Println("Timeout! No response") // 2秒後にタイムアウト
}
```

### **タイムアウトの注意点**

### 1. `time.After` のゴルーチンによるメモリリーク

- `time.After` は内部的に 新しいゴルーチンを作成する ため、大量に使うとメモリリークのリスクがある。

```go
for i := 0; i < 1000; i++ {
    select {
    case <-time.After(1 * time.Minute): // 1分ごとにタイムアウト
        fmt.Println("Timeout!")
    }
}
```

- 上記のコードは `time.After` が作るゴルーチンを解放できず、メモリを圧迫する可能性がある。
- 対策: `time.NewTimer` を使って `Stop()` でゴルーチンを明示的に解放する。

```go
timer := time.NewTimer(1 * time.Minute)

select {
case <-timer.C:
    fmt.Println("Timeout!")
}
timer.Stop() // 明示的に解放
```