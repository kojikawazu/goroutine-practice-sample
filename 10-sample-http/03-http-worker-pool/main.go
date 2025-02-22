package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Worker関数
func worker(id int, urls <-chan string, results chan<- string, wg *sync.WaitGroup) {
	// 完了通知
	defer wg.Done()

	for url := range urls {
		// リクエスト開始時間
		start := time.Now()

		// HTTPリクエストを送信
		resp, err := http.Get(url)
		if err != nil {
			results <- fmt.Sprintf("Worker %d: Failed to fetch %s: %v", id, url, err)
			continue
		}
		defer resp.Body.Close()

		results <- fmt.Sprintf("Worker %d: Fetched %s in %v", id, url, time.Since(start))
	}
}

func main() {
	// リクエストURL
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/4",
		"https://jsonplaceholder.typicode.com/posts/5",
	}

	// 3つのWorkerで並列実行
	numWorkers := 3
	urlChan := make(chan string, len(urls))
	results := make(chan string, len(urls))
	var wg sync.WaitGroup

	// Workerを起動
	for i := 1; i <= numWorkers; i++ {
		// 完了通知
		wg.Add(1)
		// Workerを起動
		go worker(i, urlChan, results, &wg)
	}

	// タスクをチャネルに送信
	for _, url := range urls {
		urlChan <- url
	}
	close(urlChan) // 送信完了を通知

	// Workerの終了を待つ
	wg.Wait()
	close(results)

	// 結果を出力
	for result := range results {
		fmt.Println(result)
	}

	// すべてのリクエストが完了したことを通知
	fmt.Println("All requests finished")
}
