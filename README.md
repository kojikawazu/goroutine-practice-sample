# Goroutine Practice Sample

## Summary

Goroutine 学習用のリポジトリです。

## Goroutine とは

Goroutine は `Go` 言語の軽量な並行処理のための機能です。  
通常のスレッドよりも軽量で、数千以上の `Goroutine` を並行に実行することが可能です。  
`Goroutine` は `OS` のスレッドではなく、`Go` のランタイムが管理する仮想スレッド上で動作します。

### **スレッドとの違い**
- `OS` スレッドは重く、コンテキストスイッチ（切り替え）にコストがかかる。
- `Goroutine` は `Go` のランタイムによって管理され、スケジューラにより効率的に `OS` スレッドにマッピングされる。


## Page

- [Goroutine の基本](./01-sample-goroutine/README.md)
- [チャネルの基本](./02-sample-channel/README.md)
- [チャネルのバッファ](./03-sample-channel-buffer/README.md)
## 実行方法

リポジトリをクローンして、各フォルダ内のサンプルコードを実行してください。

```sh
git clone https://github.com/your-repo/goroutine-practice-sample.git
cd goroutine-practice-sample
go run ./01-sample-goroutine/main.go
```
