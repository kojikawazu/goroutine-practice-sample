package main

import (
	"fmt"
	"sync"
	"time"
)

// 共有データ
var counter int

// ミューテックス
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	// タスクが完了したことをWaitGroupに通知
	defer wg.Done()

	// 5回繰り返す
	for i := 0; i < 5; i++ {
		// ロック
		mu.Lock()
		// 共有データを更新
		counter++
		// アンロック
		mu.Unlock()
		// 100ミリ秒待機
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	// WaitGroupを作成
	var wg sync.WaitGroup
	// 2つのゴルーチンを追加
	wg.Add(2)
	// ゴルーチンを起動
	go increment(&wg)
	go increment(&wg)
	// すべてのタスクが完了するまで待機
	wg.Wait()
	// 最終的なカウンターを表示
	fmt.Println("Final Counter:", counter)
}
