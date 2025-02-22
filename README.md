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
- [Select](./04-sample-select/01-select/README.md)
- [Select タイムアウト](./04-sample-select/02-select-timeout/README.md)
- [WaitGroup](./05-sample-waitgroup/README.md)
- [Mutex](./06-sample-mutex/01-mutex/README.md)
- [Mutex R/W](./06-sample-mutex/02-mutex-rw/README.md)
- [WorkerPool](./07-sample-worker-pool/README.md)
- [Context](./08-sample-context/01-context/README.md)
- [Context Timeout](./08-sample-context/02-timeout/README.md)
- [Context Deadline](./08-sample-context/03-deadline/README.md)
- [Context Value](./08-sample-context/04-value/README.md)
- [Pipeline](./09-sample-pipeline/01-simple-pipeline/README.md)
- [Pipeline Parallel](./09-sample-pipeline/02-paralell-pipeline/README.md)
- [Pipeline Context](./09-sample-pipeline/03-pipeline-context/README.md)
- [HTTP リクエスト](./10-sample-http/01-http/README.md)
- [HTTP リクエスト タイムアウト](./10-sample-http/02-http-timeout/README.md)
- [HTTP リクエスト Worker Pool](./10-sample-http/03-http-worker-pool/README.md)
- [WebSocket](./11-sample-websocket/README.md)
- [Batch](./12-sample-batch/01-batch/README.md)
- [Batch WorkerPool](./12-sample-batch/02-batch-workerpool/README.md)
- [Batch Context Cancel](./12-sample-batch/03-batch-context-cancel/README.md)

## 実行方法

リポジトリをクローンして、各フォルダ内のサンプルコードを実行してください。

```sh
git clone https://github.com/your-repo/goroutine-practice-sample.git
cd goroutine-practice-sample
go run ./01-sample-goroutine/main.go
```
