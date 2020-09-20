package main

import "fmt"

func goroutine1(s []int, c chan int){
	sum := 0
	for _, v := range s{
		sum += v
	}
	c <- sum // channelに送信
}

func goroutine2(s []int, c chan int){
	sum := 0
	for _, v := range s{
		sum += v
	}
	c <- sum // channelに送信
}

func main(){
	s := []int{1,2,3,4,5}
	c := make(chan int)
	go goroutine1(s, c)
	go goroutine2(s, c)
	x := <- c // channelを受信 blockingしている（受信するまでコードが進まない　→ sync/awaitしなくて良い）
	fmt.Println(x) // 15
	y := <- c // どんどん同じchannel経由で受信できる　もちろん別々のchannelを走らせても良い
	fmt.Println(y) // 15
}