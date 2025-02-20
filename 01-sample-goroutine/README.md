# Goroutine Sample

## Summary

Goroutine の基本系です。

## ゴルーチンとは

ゴルーチンは関数を `go` キーワードを使って呼び出すことで、並行処理を実現します。
以下のコードは、無名関数をゴルーチンとして実行し、`Hello, World!` を出力します。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    go func() {
        fmt.Println("Hello, World!")
    }()

    // main関数が即終了するとゴルーチンも終了してしまう
    time.Sleep(time.Second) // 1秒待ってから終了
}
```

## ゴルーチンの注意点

ゴルーチンは非同期で実行されるため、`main`関数が終了するとゴルーチンも終了します。

```go
package main

import (
    "fmt"
)

func main() {
    go func() {
        fmt.Println("Hello, World!")
    }()
    
    // main関数が終了するため、出力されない可能性がある
}
```

