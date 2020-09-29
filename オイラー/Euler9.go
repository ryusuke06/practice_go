/*
ピタゴラス数(ピタゴラスの定理を満たす自然数)とは a < b < c で以下の式を満たす数の組である.
a2 + b2 = c2
例えば, 32 + 42 = 9 + 16 = 25 = 52 である.

a + b + c = 1000 となるピタゴラスの三つ組が一つだけ存在する.
これらの積 abc を計算しなさい.
*/

package main

import "fmt"

func test_triplet(a int, b int, c int) bool {
	return (a*a + b*b) == c*c
}

func main() {
	a, b, c := 0, 0, 0
	for a = 1; a < 1000; a++ {
		for b = 1; b < 1000; b++ {
			for c = 1; c < 1000; c++ {
				if test_triplet(a, b, c) {
					fmt.Printf("Found triplet: %d, %d, %d (sum=%d)\n", a, b, c, a+b+c)
					if a+b+c == 1000 {
						fmt.Printf("%d^2 + %d^2 = %d^2 (%d + %d = %d)\n", a, b, c, a*a, b*b, c*c)
						fmt.Printf("%d x %d x %d = %d\n", a, b, c, a*b*c)
						return
					}
				}
			}
		}
	}
}

// A.200 x 375 x 425 = 31875000