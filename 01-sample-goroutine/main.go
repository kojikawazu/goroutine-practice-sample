package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello from goroutine")
}

func main() {
	for i := 0; i < 5; i++ {
		// ゴルーチンを起動。別のスレッドで実行される
		go hello()
	}
	fmt.Println("Hello from main")

	// ゴルーチンが終了するまで待つ
	time.Sleep(time.Second)
}
