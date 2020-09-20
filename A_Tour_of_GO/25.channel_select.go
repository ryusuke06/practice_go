package main

import (
	"fmt"
	"time"
)

func goroutine1(ch chan string){
	for {
		ch <- "packet from 1"
		time.Sleep(1 * time.Second) // 次のpacketが来るまでに1s遅れ想定
	}
}

func goroutine2(ch chan int){
	for {
		ch <- 100
		time.Sleep(3 * time.Second) // 3s遅れ想定
	}
}

func main(){
	c1 := make(chan string)
	c2 := make(chan int) // 別々のchannelなのでデータ型が異なっても問題ない
	go goroutine1(c1)
	go goroutine2(c2)

	for {
		select {
		case msg1 := <-c1:
			// 例えばc2の受信が遅れていたとしても別々のchannelを通しているのでc1はそのまま進む
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}