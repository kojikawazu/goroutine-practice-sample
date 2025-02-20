# Select Sample

## **Selectとは？**

- `select` は Go 言語の並行処理のための機能です。  
- **複数のチャネルの受信を待ち、最初にデータが到着したチャネルの処理を実行します。**  
- 通常、複数の Goroutine からのデータを受け取る場合に役立ちます。

## **Selectの使い方**

### **基本的な書き方**

- `select` は `switch` 文に似ていますが、**チャネルの送受信に対して動作する** という違いがあります。

```go
select {
case msg := <-ch1:
    fmt.Println("Received:", msg) // ch1 から受信
case msg := <-ch2:
    fmt.Println("Received:", msg) // ch2 から受信
}
```

どちらかのチャネルにデータが来たら、その処理を実行します。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // チャネルを作成
    ch1 := make(chan string)
    ch2 := make(chan string)

    // Goroutine 1: 1秒後に送信
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "From channel 1"
    }()

    // Goroutine 2: 2秒後に送信
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "From channel 2"
    }()

    // selectでどちらかのデータを受信
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1) // ch1 からのメッセージを受信
        case msg2 := <-ch2:
            fmt.Println(msg2) // ch2 からのメッセージを受信
        }
    }

    fmt.Println("Finished")
}
```

## **Selectの注意点**

### 1. default 節を使うとブロックを防げる

- 通常の `select` はどのチャネルにもデータがない場合、処理がブロックされます。
- しかし `default` 節を使うことで、即座に別の処理に進めることができます。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)

    go func() {
        time.Sleep(2 * time.Second)
        ch <- "Hello"
    }()

    for {
        select {
        case msg := <-ch:
            fmt.Println("Received:", msg)
            return
        default:
            fmt.Println("Waiting for message...")
            time.Sleep(500 * time.Millisecond) // 他の処理を実行
        }
    }
}
```