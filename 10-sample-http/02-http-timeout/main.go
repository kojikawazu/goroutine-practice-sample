package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// タイムアウト付きHTTPリクエスト
func fetchURL(ctx context.Context, url string, wg *sync.WaitGroup) {
	// 完了通知
	defer wg.Done()

	// タイムアウトを設定
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Printf("Failed to create request for %s: %v\n", url, err)
		return
	}

	// リクエスト開始時間
	start := time.Now()

	// HTTPリクエストを送信
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Failed to fetch %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Fetched %s in %v\n", url, time.Since(start))
}

func main() {
	// リクエストURL
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
	}

	// 2秒のタイムアウトを設定
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 完了通知
	var wg sync.WaitGroup

	// リクエスト送信
	for _, url := range urls {
		// 完了通知
		wg.Add(1)
		// リクエスト送信
		go fetchURL(ctx, url, &wg)
	}

	// すべてのリクエストが完了するのを待つ
	wg.Wait()
	// すべてのリクエストが完了したことを通知
	fmt.Println("All requests finished")
}
