package main

import (
	"fmt"
	"sync"
	"time"
)

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done() // 並列処理が完了したことを宣言
	for i := 0; i<5; i++{
		time.Sleep(100 * time.Millisecond)
		// timeを付けずにrunすると先にnormal()が完了、出力されgoroutine()が出力されないまま終わる
		// goroutineの出力を待ってもらうには？ → *sync.WaitGroupを使う
		fmt.Println(s)
	}
}

func normal(s string)  {
	for i := 0; i<5; i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1) // 1つ並列処理があると宣言
	go goroutine("world", &wg) // go付ければgoroutine(並列処理)できる
	normal("hello")
	// time.Sleep(2000 * time.Millisecond) 一応wg使わなくてもコードを終了させず２秒間待つなどする方法もある
	wg.Wait() // 並列処理が完了するまで待つと宣言
}