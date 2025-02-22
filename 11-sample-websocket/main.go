package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// サーバーの設定
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// クライアントの情報
type Client struct {
	conn *websocket.Conn
	id   int
}

// クライアントの情報を管理するマップ
var clients = make(map[int]*Client)

// クライアントのID
var clientID int

// 排他制御
var mu sync.Mutex

// クライアントの接続を処理する関数
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// クライアントの接続を処理
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade:", err)
		return
	}

	// クライアントのIDを管理
	mu.Lock()
	clientID++
	id := clientID
	clients[id] = &Client{conn, id}
	mu.Unlock()

	fmt.Printf("Client %d connected\n", id)

	// クライアントの切断を処理
	defer func() {
		mu.Lock()
		delete(clients, id)
		mu.Unlock()
		conn.Close()
		fmt.Printf("Client %d disconnected\n", id)
	}()

	// クライアントからのメッセージを受信
	for {
		// クライアントからのメッセージを受信
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Client %d connection closed: %v\n", id, err)
			break
		}
		fmt.Printf("Client %d sent: %s\n", id, string(msg))

		// クライアント全員にメッセージをブロードキャスト
		broadcast(msg)
	}
}

// メッセージを全クライアントにブロードキャストする関数
func broadcast(msg []byte) {
	// 排他制御
	mu.Lock()
	defer mu.Unlock()

	// 全クライアントにメッセージを送信
	for _, client := range clients {
		err := client.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Printf("Failed to send message to Client %d: %v\n", client.id, err)
		}
	}
}

// サーバーの起動
func main() {
	// サーバーの起動
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("WebSocket server started on ws://localhost:8080/ws")
	http.ListenAndServe(":8080", nil)
}
