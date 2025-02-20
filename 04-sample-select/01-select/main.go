package main

import (
	"fmt"
	"time"
)

func main() {
	// チャネルを作成
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1: 1秒後に送信
	// 無名関数も設定可能
	go func() {
		time.Sleep(time.Second)
		ch1 <- "From channel 1"
	}()

	// Goroutine 2: 2秒後に送信
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "From channel 2"
	}()

	// selectでどちらかのデータを受信
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1) // ch1 からのメッセージを受信
		case msg2 := <-ch2:
			fmt.Println(msg2) // ch2 からのメッセージを受信
		}
	}

	// すべてのメッセージを受信したことを表示
	fmt.Println("Finished")
}
