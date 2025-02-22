package main

import (
	"fmt"
)

// ステージ1: データ生成（Producer）
func generate(numbers ...int) <-chan int {
	// チャネルを作成
	out := make(chan int)

	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out) // 送信完了を通知
	}()

	// チャネルを返す
	return out
}

// ステージ2: データ処理（Worker）
func multiplyByTwo(in <-chan int) <-chan int {
	// チャネルを作成
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out) // 送信完了を通知
	}()

	// チャネルを返す
	return out
}

func main() {
	// ステージ1: データ生成
	numbers := generate(1, 2, 3, 4, 5)

	// ステージ2: データ変換
	results := multiplyByTwo(numbers)

	// ステージ3: データ収集（Consumer）
	for result := range results {
		fmt.Println(result)
	}
}
