package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second) // ２秒間もかかる関数と想定
	fmt.Println("finish")
	ch <- "result"
}

func main()  {
	ch := make(chan string)
	ctx := context.Background() // contextを作る
	ctx, cancel := context.WithTimeout(ctx, 3 * time.Second)
	// 3秒間かかってgoroutineが終わらなければDone()
	// contextを引数にとるがタイムアウトとかを使わなければとりあえず
	// ctx := context.TODO() 何をするか決めてないという意味でTODO()を使う
	defer cancel()
	go longProcess(ctx, ch)

	CTXLOOP:
		for {
			select {
			case <- ctx.Done():
				fmt.Println(ctx.Err())
				break CTXLOOP
			case <- ch:
				fmt.Println("success")
				break CTXLOOP
			}
		}
	fmt.Println("#####")
}