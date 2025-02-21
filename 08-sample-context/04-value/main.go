package main

import (
	"context"
	"fmt"
)

func worker(ctx context.Context) {
	userID := ctx.Value("userID") // コンテキストから値を取得
	fmt.Println("Worker started for user:", userID)
}

func main() {
	ctx := context.WithValue(context.Background(), "userID", 12345)

	go worker(ctx)

	fmt.Println("Main function finished.")
}
