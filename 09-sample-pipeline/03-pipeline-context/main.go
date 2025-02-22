package main

import (
	"context"
	"fmt"
	"time"
)

// データ生成
func generate(ctx context.Context, numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range numbers {
			select {
			case out <- n:
			case <-ctx.Done():
				fmt.Println("Generate canceled")
				return
			}
		}
	}()
	return out
}

// データ処理
func multiplyByTwo(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * 2:
			case <-ctx.Done():
				fmt.Println("Multiply canceled")
				return
			}
		}
	}()
	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	numbers := generate(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	results := multiplyByTwo(ctx, numbers)

	for result := range results {
		fmt.Println(result)
		time.Sleep(500 * time.Millisecond) // 遅延をシミュレート
	}
}
