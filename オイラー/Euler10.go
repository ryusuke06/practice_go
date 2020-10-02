package main

import "fmt"

func main() {
	var n int // 答え
	j := []int{2} // 素数をまとめるスライス

	for i := 3;i <= 2000000;i++ {
		if i%2!=0{ // 2で割り切れる物は弾く=>無駄にforループしないで済む
		LOOP:
			for x := 0;x >= 0;x++ {
				if len(j)==x { // jにある素数の全部で割っても割り切れなかったら=>素数
					j = append(j, i) // Unshift
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

// 142913828920