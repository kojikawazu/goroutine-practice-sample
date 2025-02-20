package main

import (
	"fmt"
	"time"
)

// workerはチャネルを受け取り、メッセージを送信する
func worker(id int, ch chan string) {
	for i := 0; i < 3; i++ {
		// 1秒待ってからメッセージを送信
		time.Sleep(time.Second)
		ch <- fmt.Sprintf("Worker %d: Task %d: completed", id, i+1)
	}
}

func main() {
	// チャネルを作成
	ch := make(chan string)

	// 3つのworkerを起動
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	// 6回チャネルからメッセージを受け取る
	for i := 0; i < 6; i++ {
		msg := <-ch
		fmt.Println(msg)
	}

	// すべてのタスクが完了したことを表示
	fmt.Println("All tasks completed")
}
