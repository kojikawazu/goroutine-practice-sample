package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // タイムアウト時にキャンセル
			fmt.Println("Worker stopped due to timeout")
			return
		default:
			fmt.Println("Worker is working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // メモリリーク防止のため必ず `defer cancel()`

	go worker(ctx)

	time.Sleep(3 * time.Second) // タイムアウト後も待つ
	fmt.Println("Main function finished.")
}
