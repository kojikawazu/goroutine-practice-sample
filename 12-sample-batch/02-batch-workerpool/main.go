package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker 関数
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	// ジョブが完了したら WaitGroup のカウントを確実に減らす
	defer wg.Done()

	// チャネルからジョブを受け取り、処理を行う
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // 擬似的な処理時間
		results <- fmt.Sprintf("Job %d processed by Worker %d", job, id)
	}
}

func main() {
	// ジョブの数
	numJobs := 10
	// 並列実行するWorkerの数
	numWorkers := 3

	// タスク用チャネル
	jobs := make(chan int, numJobs)
	// 結果用チャネル
	results := make(chan string, numJobs)
	// WaitGroup
	var wg sync.WaitGroup

	// Worker を起動（固定数のGoroutine）
	for i := 1; i <= numWorkers; i++ {
		// ジョブが完了したら WaitGroup のカウントを確実に減らす
		wg.Add(1)
		// Worker を起動
		go worker(i, jobs, results, &wg)
	}

	// タスクをチャネルに送信
	for j := 1; j <= numJobs; j++ {
		// タスクをチャネルに送信
		jobs <- j
	}

	// Worker の完了を待つ
	wg.Wait()
	// 結果用チャネルをクローズ
	close(results)

	// 結果を出力
	for result := range results {
		fmt.Println(result)
	}

	// すべてのジョブが完了したことを出力
	fmt.Println("All batch jobs completed!")
}
