# Channel Buffer Sample

## **バッファとは？**

- バッファとは、Go 言語のチャネルに一定数のデータを蓄える機能です。  
- 通常のチャネル（バッファなし）はデータを送信すると、すぐに受信しないとブロックされますが、バッファ付きチャネルを使うことで、一定数のデータをチャネル内に一時保存できます。

### **バッファあり vs バッファなし**

| 種類 | 送信時 | 受信時 |
|------|------|------|
| **バッファなし** (`make(chan string)`) | 受信側がデータを受け取るまでブロック | データが来るまでブロック |
| **バッファあり** (`make(chan string, 5)`) | バッファに空きがあれば送信できる | バッファが空ならブロック |

## **バッファの使い方**

バッファは `make(chan type, バッファサイズ)` のように作成します。

```go
ch := make(chan string, 5) // バッファサイズ5のチャネルを作成
```

バッファにデータを送信すると、バッファが埋まるまでは即座に送信されます。

```go
package main

import "fmt"

func main() {
    ch := make(chan string, 3) // バッファサイズ3

    ch <- "A" // 即座に送信
    ch <- "B" // 即座に送信
    ch <- "C" // 即座に送信

    fmt.Println(<-ch) // "A"
    fmt.Println(<-ch) // "B"
    fmt.Println(<-ch) // "C"
}
```

## **バッファの注意点**

### 1. main関数が終了するとバッファも終了する

バッファ付きチャネルも main 関数が終了するとプログラムが停止します。

```go
package main

import "fmt"

func main() {
    ch := make(chan string, 5)
    
    go func() {
        ch <- "Hello, Buffer!"
    }()

    // `main` が終了するため、メッセージを受け取れない可能性あり
}
```

### 2. バッファが満杯のときは送信がブロックされる

```go
package main

import "fmt"

func main() {
    ch := make(chan string, 2) // バッファサイズ2

    ch <- "A"
    ch <- "B"

    // ch <- "C" // ここでブロックしてしまう！（バッファが満杯）
    fmt.Println(<-ch)
    fmt.Println(<-ch)

    ch <- "C" // ここならOK
    fmt.Println(<-ch)
}
```