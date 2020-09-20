package main

import "fmt"

func goroutine(s []int, c chan int){
	sum := 0
	for _, v := range s{
		sum += v
		c <- sum
	}
	close(c) // channelに順繰りに全て送信したら受診しようとしている対象に対して終わるように宣言
}

func main(){
	s := []int{1,2,3,4,5}
	c := make(chan int) // c := make(chan int, len(s))でも可 channelのメモリーを確保したいなら
	go goroutine(s, c)
	for i := range c { // close()を受信するまで随時更新
		fmt.Println(i)
	}
}