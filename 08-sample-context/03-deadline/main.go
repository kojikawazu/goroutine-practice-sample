package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped due to deadline")
			return
		default:
			fmt.Println("Worker is working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	deadline := time.Now().Add(2 * time.Second) // 2秒後の時間を指定
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go worker(ctx)

	time.Sleep(3 * time.Second) // タイムアウト後も待機
	fmt.Println("Main function finished.")
}
