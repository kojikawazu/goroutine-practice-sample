package main

import (
	"fmt"
	"sync"
)

// 共有データ
var counter int

// 読み書きロック
var rwMu sync.RWMutex

// readCounterはカウンターを読み取る
func readCounter(wg *sync.WaitGroup) {
	// タスクが完了したことをWaitGroupに通知
	defer wg.Done()

	// 読み取りロック
	rwMu.RLock()
	// カウンターを読み取る
	fmt.Println("Read Counter:", counter)
	// 読み取りロック解除
	rwMu.RUnlock()
}

// writeCounterはカウンターを書き込む
func writeCounter(wg *sync.WaitGroup) {
	// タスクが完了したことをWaitGroupに通知
	defer wg.Done()

	// 書き込みロック
	rwMu.Lock()
	// カウンターを書き込む
	counter++
	// 書き込みロック解除
	rwMu.Unlock()
}

func main() {
	// WaitGroupを作成
	var wg sync.WaitGroup

	// 5回繰り返す
	for i := 0; i < 5; i++ {
		// WaitGroupに読み取りGoroutineを追加
		wg.Add(1)
		// 読み取りGoroutineを実行
		go readCounter(&wg)
	}

	// 2回繰り返す
	for i := 0; i < 2; i++ {
		// WaitGroupに書き込みGoroutineを追加
		wg.Add(1)
		// 書き込みGoroutineを実行
		go writeCounter(&wg)
	}

	// すべてのタスクが完了するまで待機
	wg.Wait()
	// 最終的なカウンターを表示
	fmt.Println("Final Counter:", counter)
}
