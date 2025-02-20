# WaitGroup Sample

## WaitGroupとは？

- `sync.WaitGroup` は Go 言語の並行処理で、**複数の Goroutine の終了を待つための機能** です。
- 通常、Goroutine は非同期で動作するため、`main` 関数が終了するとすべての Goroutine も終了してしまいます。
- `sync.WaitGroup` を使うことで、すべての Goroutine が完了するのを確実に待機できます。

## WaitGroupの使い方

基本的な使い方は以下の 3 ステップです。

1. `wg.Add(n)`  
   - 実行する Goroutine の数を指定
2. `wg.Done()`  
   - Goroutine の処理が完了したことを通知（カウントを減らす）
3. `wg.Wait()`  
   - すべての Goroutine が完了するまで待機

### 基本的なサンプル

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Goroutine の終了を通知
    fmt.Printf("Worker %d started\n", id)
    time.Sleep(time.Second) // 処理のシミュレーション
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // Goroutine の数を追加
        go worker(i, &wg)
    }

    wg.Wait() // すべての Goroutine が終了するのを待つ
    fmt.Println("All workers finished")
}
```

## WaitGroupの注意点

### 1. `wg.Add(n)` を `go` の前に呼ぶ

- 以下のコードでは `wg.Add(1)` を `go` の前に呼ぶ必要があります。
- `Add()` を `go` の後に呼ぶと、Goroutine の実行順によっては `Wait()` が先に実行され、Goroutine が待機されない可能性があります。

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d started\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // `go` の前に呼ぶ
        go worker(i, &wg)
    }

    wg.Wait()
    fmt.Println("All workers finished")
}
```

### 2. `wg.Wait()` はカウンタが 0 の場合、即時終了

- `wg.Wait()` はカウンタが 0 の場合、即座に終了します。
- もし `wg.Add(n)` をせずに `Wait()` すると、何も待たずに次の処理に進みます。

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Wait() // すぐに終了する（Goroutine が 0 なので）
    fmt.Println("Done!")
}
```

### 3. `wg.Done()` を忘れるとデッドロック

- `wg.Done()` を呼び忘れると、`wg.Wait()` がずっと待機状態になり、デッドロックします。

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    fmt.Printf("Worker %d started\n", id)
    time.Sleep(time.Second)
    // wg.Done() を忘れると、wg.Wait() が無限待機してしまう！
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait() // ここで無限待機（Done が呼ばれていない）
    fmt.Println("All workers finished")
}
```