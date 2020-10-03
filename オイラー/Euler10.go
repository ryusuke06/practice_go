/*
10以下の素数の和は 2 + 3 + 5 + 7 = 17 である.

200万以下の全ての素数の和を求めよ.
*/

package main

import "fmt"

func main() {
	n := 2 // 答え
	j := []int{2} // 素数をまとめるスライス

	for i := 3;i <= 2000000;i++ {
		if i%2!=0{ // 2で割り切れる物は弾く=>無駄にforループしないで済む(2は既に入れてるし)
		LOOP:
			for x := 0;x >= 0;x++ {
				if len(j)==x { // jにある素数の全部で割っても割り切れなかったら=>素数
					j = append(j, i) // Push (Unshift(前から値を追加)したいときはjをrangeで回してappend([]int{i},j[y])なんてしたらいいんでね？)
					n += i
					break LOOP
				}else if i%j[x]==0 { // 割れたら素数ではない
					break LOOP
				}
			}
		}
	}
	fmt.Println(n)
}

// 142913828922