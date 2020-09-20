package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond) // tickとafterはchannelを介して来ている
	boom := time.After(500 * time.Millisecond)
	OuterLoop:
		for {
			select {
			case <-tick:
			// tickには時間などが含まれているが必要なければ<-tickでいい（必要ならt := <-tickとか書けばいい）
				fmt.Println("tick.")
			case <-boom:
				fmt.Println("BOOM!")
				break OuterLoop // selectにbreak文があるわけではないのでラベルで書く必要あり
				// return boomが来たらmain()を終了 → for文以降のコードが走らずにプログラム終了
			default: // 2つのchannel以外から来たものを扱うのがdefault selection
				fmt.Println("    .")
				time.Sleep(50 * time.Millisecond)
			}
		}
	fmt.Println("returnだと出てこないけどlabelでbreakしとけば出て来るよ")
}
