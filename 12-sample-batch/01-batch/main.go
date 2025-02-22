package main

import (
	"fmt"
	"sync"
	"time"
)

// バッチ処理（並列実行）
func processBatch(id, value int, wg *sync.WaitGroup) {
	// 処理が終わったらDoneを呼び出す
	defer wg.Done()

	// 処理中のワーカーのIDと値を表示
	fmt.Printf("Worker %d: Processing value %d\n", id, value)
	time.Sleep(time.Millisecond * 500) // 擬似的な処理時間

	// 処理が終わったら、ワーカーのIDと値を表示
	fmt.Printf("Worker %d: Finished processing value %d\n", id, value)
}

func main() {
	// 処理する値
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 処理が終わったらDoneを呼び出すためのWaitGroup
	var wg sync.WaitGroup

	// 並列処理
	for i, v := range values {
		// 処理が終わったらDoneを呼び出すためのWaitGroupに1を追加
		wg.Add(1)
		// 処理を並列に実行
		go processBatch(i+1, v, &wg)
	}

	// すべての処理が終わるのを待つ
	wg.Wait()

	// すべての処理が終わったら、完了を表示
	fmt.Println("All batch jobs completed!")
}
