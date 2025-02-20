package main

import (
	"fmt"
	"sync"
	"time"
)

// workerはWaitGroupを受け取り、タスクを実行する
func worker(id int, duration time.Duration, wg *sync.WaitGroup) {
	// タスクが完了したことをWaitGroupに通知
	defer wg.Done()
	// タスク開始
	fmt.Printf("Worker %d started (duration: %v)\n", id, duration)
	// タスク実行
	time.Sleep(duration)
	// タスク完了
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	// WaitGroupを作成
	var wg sync.WaitGroup
	// タスクを作成
	tasks := []time.Duration{3 * time.Second, 1 * time.Second, 2 * time.Second}

	// タスクを並行実行
	for i, duration := range tasks {
		// WaitGroupにタスクを追加
		wg.Add(1)
		// タスクを並行実行
		go worker(i+1, duration, &wg)
	}

	// すべてのタスクが完了するまで待機
	wg.Wait()
	// すべてのタスクが完了したことを表示
	fmt.Println("All workers finished")
}
