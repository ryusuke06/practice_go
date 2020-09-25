/*
素数を小さい方から6つ並べると 2, 3, 5, 7, 11, 13 であり, 6番目の素数は 13 である.

10 001 番目の素数を求めよ.
*/

package main

import "fmt"

func main() {
	var n = [10001]int{2} // 固定長配列にしたのは余計なメモリ確保を避けるため＆明示的に示したかった
	for i := 3;i > 0;i++ {
		if i%2!=0{
			LOOP:
			for x := 0;x < 10001;x++ { // indexは0も含めるのでx < 10001
				if n[x]==0 { // 先に書いておかないと0で割ろうとしてしまう=>err
					n[x] = i
					fmt.Println(n[x])
					break LOOP
				}else if i%n[x]==0 {
					break LOOP
				}
			}
		}
	}
}

// A.104743