package main

import (
	"fmt"
	"sync"
	"time"
)

// workerはタスクを受け取り、結果を送信する
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	// タスクが完了したことをWaitGroupに通知
	defer wg.Done()

	// タスクを受け取る
	for job := range jobs {
		// タスクを表示
		fmt.Printf("Worker %d started job %d\n", id, job)
		// 擬似的な処理時間
		time.Sleep(time.Second)
		// 結果を送信
		results <- fmt.Sprintf("Job %d done by Worker %d", job, id)
	}
}

func main() {
	// タスク数
	numJobs := 10
	// ワーカー数
	numWorkers := 3

	// タスクを格納するChannel
	jobs := make(chan int, numJobs)
	// 処理結果を格納するChannel
	results := make(chan string, numJobs)
	// WaitGroup
	var wg sync.WaitGroup

	// Workerを起動
	for i := 1; i <= numWorkers; i++ {
		// WaitGroupにWorkerを追加
		wg.Add(1)
		// Workerを起動
		go worker(i, jobs, results, &wg)
	}

	// タスクを送信
	for j := 1; j <= numJobs; j++ {
		// タスクを送信
		jobs <- j
	}
	// 送信完了を通知
	close(jobs)

	// Workerの終了を待つ
	wg.Wait()
	// 結果を出力
	for result := range results {
		fmt.Println(result)
	}

	// すべてのタスクが完了したことを表示
	fmt.Println("All jobs are done!")
}
