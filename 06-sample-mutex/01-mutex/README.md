# Mutex Sample

## Mutexとは？

- `sync.Mutex` は **Go 言語で並行処理を安全に行うための仕組み** です。  
- 複数の Goroutine が同じデータにアクセスすると **データ競合（Race Condition）** が発生することがあります。  
- `sync.Mutex` を使うことで、ある Goroutine がデータを操作している間、他の Goroutine のアクセスを防ぐことができます。

## Mutexの使い方

### 基本的な書き方

```go
var mu sync.Mutex // ミューテックスの宣言

mu.Lock()   // ロックを取得（他の Goroutine を待機させる）
counter++   // 共有データを更新
mu.Unlock() // ロックを解除（他の Goroutine が処理可能に）
```

更に `defer mu.Unlock()` を使うことで、ロック解除忘れを防ぐのが一般的です。

```go
mu.Lock()
defer mu.Unlock()
counter++
```

## Mutexの注意点

### 1. `defer mu.Unlock()` を推奨

- ロックを取得したら、必ず `Unlock()` する必要があります。
- `defer` を使うことで、関数が途中で終了してもロック解除が確実に行われます。

```go
mu.Lock()
defer mu.Unlock()
counter++
```

### 2. デッドロックの回避

- デッドロック（Deadlock） とは、複数の `Goroutine` が互いのロック解除を待ち続ける状態のことです。

デッドロックが発生するコード

```go
package main

import (
    "fmt"
    "sync"
)

var mu1, mu2 sync.Mutex

func process1() {
    mu1.Lock()
    fmt.Println("process1: locked mu1")

    mu2.Lock() // ここで process2 が mu2 を持っているとデッドロック！
    fmt.Println("process1: locked mu2")

    mu2.Unlock()
    mu1.Unlock()
}

func process2() {
    mu2.Lock()
    fmt.Println("process2: locked mu2")

    mu1.Lock() // ここで process1 が mu1 を持っているとデッドロック！
    fmt.Println("process2: locked mu1")

    mu1.Unlock()
    mu2.Unlock()
}

func main() {
    go process1()
    go process2()

    fmt.Scanln() // Enterキーを押すまで待機
}
```

### デッドロックの回避策

- ロックの順序を統一する（常に mu1 → mu2 の順でロックする）。
- `sync.RWMutex`（読み書きロック）を活用する（読み取りはロック不要にする）。

```go
var mu1, mu2 sync.Mutex

// ロックの順序を統一
func process1() {
    mu1.Lock()
    defer mu1.Unlock()
    mu2.Lock()
    defer mu2.Unlock()
}
```