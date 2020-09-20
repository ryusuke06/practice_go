package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int, i int) {
	// Something
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup){
	for i := range ch{
		// main()でclose()しなければerr、goroutine内ならばchに次がこないか待ってる状態になる
		// だからmain()でclose()する
		func (){
			defer wg.Done()
			// 何らかの理由でプロセスがうまくいかなくてもDone()が呼ばれるように
			// inner functionでdeferを使うなんて方法もあるよ！！
			fmt.Println("process", i * 1000)
		}()
	}
	fmt.Println("###")
	// main()でclose()しないとforが次待ちのまま、Wait()が終わり
	// この"###"が出力されないままプログラムが終わる
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Producer
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i) // 10回channelに送信
	}

	// Consumer
	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
	time.Sleep(2 * time.Second)
}