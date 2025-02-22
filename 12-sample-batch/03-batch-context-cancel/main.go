package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Worker 関数
func worker(ctx context.Context, id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	// ジョブが完了したら WaitGroup のカウントを確実に減らす
	defer wg.Done()

	for {
		select {
		case <-ctx.Done(): // キャンセル信号を受信したら終了
			fmt.Printf("Worker %d: Cancelled\n", id)
			return
		case job, ok := <-jobs:
			// チャネルが閉じられたら終了
			if !ok {
				return
			}
			// ジョブを処理
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(time.Second)
			// 結果をチャネルに送信
			results <- fmt.Sprintf("Job %d processed by Worker %d", job, id)
		}
	}
}

func main() {
	// ジョブの数
	numJobs := 10
	// 並列実行するWorkerの数
	numWorkers := 3
	// キャンセル可能なコンテキストを作成
	ctx, cancel := context.WithCancel(context.Background())
	// WaitGroup
	var wg sync.WaitGroup
	// タスク用チャネル
	jobs := make(chan int, numJobs)
	// 結果用チャネル
	results := make(chan string, numJobs)

	// Worker を起動
	for i := 1; i <= numWorkers; i++ {
		// ジョブが完了したら WaitGroup のカウントを確実に減らす
		wg.Add(1)
		// Worker を起動
		go worker(ctx, i, jobs, results, &wg)
	}

	// タスクを送信
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 途中でキャンセル
	time.Sleep(3 * time.Second)
	cancel() // 全 Worker を停止

	// Worker の終了を待つ
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
