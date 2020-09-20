package main

import "fmt"

/*
fan-inは、複数の入力を1つのchannelにまとめて受信するパターンです。
fan-outは、複数の関数（goroutine）が、1つのchannelから値を読み取り送信するパターンです。
fan-outはCPUに作業を分配するため、Workersパターンとも呼ばれます。
*/

func producer(first chan int){
	defer close(first)
	for i := 0; i < 10; i++{
		first <- i
	}
}

func multi2(first chan int, second chan int){
	//   first <-chan int, second chan<- int とchannelの流れを明示的に書くこともできる
	defer close(second)
	for i := range first{
		second <- i * 2
	}
}

func multi4(second chan int, third chan int){
	defer close(third)
	for i := range second{
		third <- i * 4
	}
}

func main(){
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

	go producer(first) // ex) 1
	go multi2(first, second) // ex) 1 * 2 = 2
	go multi4(second, third) // ex) 2 * 4 = 8
	for result := range third{
		fmt.Println(result) // ex) 8 このようにステージごとに出力される(次のステージは2からなので結果は16)
	}
}