package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// HTTPリクエストを並列に送信
func fetchURL(url string, wg *sync.WaitGroup) {
	// 完了通知
	defer wg.Done()

	// リクエスト開始時間
	start := time.Now()

	// HTTPリクエストを送信
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to fetch %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	// リクエスト完了時間
	fmt.Printf("Fetched %s in %v\n", url, time.Since(start))
}

func main() {
	// リクエストURL
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
	}

	var wg sync.WaitGroup

	// 並列にHTTPリクエストを送信
	for _, url := range urls {
		// 完了通知
		wg.Add(1)
		// リクエスト送信
		go fetchURL(url, &wg)
	}

	// すべてのリクエストが完了するのを待つ
	wg.Wait()

	// すべてのリクエストが完了したことを通知
	fmt.Println("All requests finished")
}
