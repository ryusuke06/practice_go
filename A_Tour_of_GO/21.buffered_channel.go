package main

import "fmt"

func main(){
	ch := make(chan int, 2) // buffered channel 二つの入り口しかないと指定
	ch <- 100
	fmt.Println(len(ch)) // 1 （入り口の一つは100が利用）
	ch <- 200
	fmt.Println(len(ch)) // 2 （入り口の一つは200が利用）
	// ch <- 300 // err （len(ch)==2なため）
	// fmt.Println(len(ch))　err 入り口をどちらかをx := <-chなどで排出してからでないと送信不可

	close(ch) // close()がないとfor文で繰り返したときlen(ch)==0でも排出しようとしてerrになる
	for c := range ch{
		fmt.Println(c)
	}
}