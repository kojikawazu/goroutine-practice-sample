package main

import (
	"fmt"
	"time"
)

func main() {
	// チャネルを作成
	ch := make(chan string)

	// 3秒後にメッセージを送信
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Hello, after delay"
	}()

	// チャネルからメッセージを受信
	select {
	case msg := <-ch:
		fmt.Println(msg)
	// 2秒後にタイムアウト
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout! No response")
	}
}
