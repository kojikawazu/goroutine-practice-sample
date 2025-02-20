# Mutex R/W Sample

## Mutex R/Wとは？

- `sync.RWMutex` は Go 言語の並行処理で **読み取りと書き込みを区別する排他制御** を提供します。  
- 通常の `sync.Mutex` では **読み取り時もロックされる** ため、読み取りが多い処理では `sync.RWMutex` を使うとパフォーマンスが向上します。

## 通常の `sync.Mutex` との違い

| 種類 | 読み取り (`RLock`) | 書き込み (`Lock`) |
|------|-----------------|----------------|
| `sync.Mutex` | **不可**（書き込みと同じロックを取得） | **不可**（他の読み書きが完了するまで待機） |
| `sync.RWMutex` | **可**（複数 Goroutine が同時に読み取り可能） | **不可**（他の読み書きが完了するまで待機） |

## 要点

- 複数の Goroutine が同時に読み取りできる
- 書き込み中は読み取りも待機する

## Mutex R/Wの使い方

### 基本的な書き方

```go
var rwMu sync.RWMutex // 読み書きロックの宣言

rwMu.RLock()   // 読み取りロック
fmt.Println("Read data")
rwMu.RUnlock() // 読み取りロック解除

rwMu.Lock()   // 書き込みロック
counter++     // 共有データを更新
rwMu.Unlock() // 書き込みロック解除
```

`defer rwMu.Unlock()` / `defer rwMu.RUnlock()` を使うことで、ロック解除忘れを防ぐのが一般的です。

## Mutex R/Wの注意点

### 1. 書き込み中は読み取りもブロックされる

- `Lock()` を使うと、読み取り Goroutine も 書き込みが終わるまで待機 する。
- `RLock()` は複数の Goroutine で同時に実行できるが、`Lock()` が呼ばれると すべての `RLock()` がブロックされる。

### 2. `defer rwMu.Unlock()` / `defer rwMu.RUnlock()` を推奨

- ロックを取得したら、必ず解除する必要があります。
- `defer` を使うことで、関数が途中で終了してもロック解除が確実に行われます。

```go
rwMu.RLock()
defer rwMu.RUnlock()
fmt.Println("Reading data")

rwMu.Lock()
defer rwMu.Unlock()
counter++
```