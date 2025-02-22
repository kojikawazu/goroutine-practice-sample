# WebSocket Server with Gorilla WebSocket

## WebSocket サーバーの実装

- `Go` の `Gorilla WebSocket` を利用して、リアルタイムな通信を可能にする WebSocket サーバーを実装します。
- クライアントを管理し、接続・切断の管理を行います。
- ブロードキャスト機能を実装し、クライアント間でメッセージを共有できるようにします。

## 処理の流れ

1. クライアントが `/ws` エンドポイントに接続すると、WebSocket 接続を確立
2. クライアントごとに `Client` インスタンスを作成し、管理用の `clients` マップに登録
3. クライアントからのメッセージを受信
4. 受信したメッセージを全クライアントにブロードキャスト
5. クライアントが切断された場合、管理マップから削除

## 実行方法

### 1. ライブラリのインストール

```bash
go get github.com/gorilla/websocket
```

### 2. WebSocket サーバーを起動

```bash
go run main.go
```

- サーバーが起動すると、以下のようなメッセージが表示されます。

```bash
WebSocket server started on ws://localhost:8080/ws
```

### 3. WebSocket クライアントに接続

- WebSocket クライアントツール（例: wscat）を使用して接続する:

```bash
wscat -c ws://localhost:8080/ws
```

- ブラウザのコンソール から WebSocket クライアントを作成:

```javascript
let ws = new WebSocket("ws://localhost:8080/ws");

ws.onmessage = function(event) {
    console.log("Received:", event.data);
};

ws.onopen = function() {
    ws.send("Hello, Server!");
};
```

### 4. クライアント間でメッセージを送受信

- クライアント 1 で送信

```bash
> Hello, Server!
```

- 他のクライアントにも同じメッセージが届く

```bash
Received: Hello, Server!
```

## メリット

- リアルタイム通信
  - サーバーとクライアント間で双方向通信が可能。
- 複数クライアント対応
  - 複数のクライアントを管理し、メッセージをブロードキャストできる。
- Goroutine を活用
  - 各クライアント接続は Goroutine で処理されるため、非同期に処理可能。

## 注意点

### 1. sync.Mutex による並行アクセス制御

- 複数の Goroutine が clients マップを更新するため、sync.Mutex を使用して排他制御。

```go
mu.Lock()
clients[id] = &Client{conn, id}
mu.Unlock()
```

### 2. CheckOrigin の設定

- デフォルトでは WebSocket の CORS がブロックされるため、開発環境では CheckOrigin: func(r *http.Request) bool { return true } を設定。

```go
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}
```

- 本番環境では CheckOrigin の設定を適切に行い、信頼できるオリジンのみ許可するようにする。

### 3. クライアントの切断処理

- クライアントが切断された際に clients マップから削除。

```go
defer func() {
    mu.Lock()
    delete(clients, id)
    mu.Unlock()
    conn.Close()
}()
```
