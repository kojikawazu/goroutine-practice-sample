# Channel Sample

## チャネルとは

- チャネルは `Go` 言語の並行処理のための機能です。  
- `Goroutine` 同士がデータをやり取りするために使用されます。  
- チャネルを利用することで、安全にデータの受け渡しができます。

## **チャネルの作成と基本操作**

チャネルは `make` 関数を使って作成します。

```go
ch := make(chan string) // string型のチャネルを作成
```

## データの送信と受信

チャネルには `<-` 演算子を使ってデータを送受信します。

```go
package main

import "fmt"

func main() {
    ch := make(chan string) // チャネル作成

    // ゴルーチン内でメッセージを送信
    go func() {
        ch <- "Hello, Channel!"
    }()

    // メイン関数でメッセージを受信
    msg := <-ch
    fmt.Println(msg) // "Hello, Channel!"
}
```

## チャネルの注意点

### 1. main関数の終了でチャネルも終了する

チャネルは非同期で動作するため、メイン関数が終了するとデータを受け取る前にプログラムが終了することがあります。

```go
package main

import "fmt"

func main() {
    ch := make(chan string)

    go func() {
        ch <- "Hello, World!"
    }()

    // ここで main 関数が即終了すると、チャネルのデータが受信されない
}
```

→ 対策: `sync.WaitGroup` や `time.Sleep` を使って `Goroutine` の完了を待つ

### 2. close() を適切に使う

チャネルは明示的に close() しない限り、受信側は無限待機状態になることがあります。

```go
package main

import "fmt"

func main() {
    ch := make(chan string)

    go func() {
        ch <- "Hello, Close!"
        close(ch) // チャネルをクローズ
    }()

    for msg := range ch { // クローズされるまで受信
        fmt.Println(msg)
    }
}
```

→ 対策: `for msg := range ch` を使ってチャネルがクローズされるまで受信する
