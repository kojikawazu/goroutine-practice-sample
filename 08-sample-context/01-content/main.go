package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done(): // 親Goroutineからキャンセル信号を受け取ったら終了
			fmt.Printf("Worker %d stopped\n", id)
			return
		default:
			fmt.Printf("Worker %d is working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Cancelling all workers...")
	cancel() // すべてのGoroutineを停止

	time.Sleep(1 * time.Second) // 終了を確認
	fmt.Println("All workers stopped.")
}
