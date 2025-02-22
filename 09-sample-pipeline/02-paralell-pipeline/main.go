package main

import (
	"fmt"
	"sync"
)

// ステージ1: データ生成
func generate(numbers ...int) <-chan int {
	// チャネルを作成
	out := make(chan int)

	// データを送信
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()

	// チャネルを返す
	return out
}

// ステージ2: 並列処理（Worker Pool）
func multiplyByTwo(in <-chan int, workerCount int) <-chan int {
	// チャネルを作成
	out := make(chan int)

	// グループを作成
	var wg sync.WaitGroup

	// 複数のWorkerを起動
	for i := 0; i < workerCount; i++ {
		// グループに追加
		wg.Add(1)

		// Workerを起動
		go func() {
			defer wg.Done()

			// データを受け取る
			for n := range in {
				out <- n * 2
			}
		}()
	}

	// 全Workerの完了後に出力チャネルを閉じる
	go func() {
		// 全Workerの完了を待つ
		wg.Wait()

		// 出力チャネルを閉じる
		close(out)
	}()

	return out
}

func main() {
	// データを生成
	numbers := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// データを2倍にする
	results := multiplyByTwo(numbers, 3) // 3つのWorkerで並列処理

	// データを受け取る
	for result := range results {
		fmt.Println(result)
	}
}
