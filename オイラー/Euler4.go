package main

import (
	"fmt"
	"strconv"
)

func is_palindrome(i int64) bool {
	s := strconv.FormatInt(i, 10) // 数値型を文字列に変換
	i, j := int64(0), int64(len(s)-1) // 文字列のindex番号を始めをi、末をjとする
	for i < j { // i++、j--していって文字が一致する確認 = 回文数であることを確認
		if s[i] != s[j] { // string[index]で...,'0':48,'1':49,...,'9':57,...と文字に対して割り当てられた数値（アスキーコード）が返ってくる。つまり s[i] != s[j] 同じ文字でないことを示す ちなみにstring()でキャストできる
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	product := int64(0)
	largest_product := int64(0)
	i, j := int64(999), int64(999)

	for i > 99 && j > 99 {
		product = i * j
		if is_palindrome(product) && product > largest_product {
			largest_product = product
			fmt.Printf("%d from %d x %d\n", product, i, j)
			i--
			j = 999
			if i*j < largest_product {
				break
			}
		} else if j > 100 { // j > 99 だと最終的に j = 99 となり、継続式から外れて次の繰り返しが行われないから j > 100
			j -= 1
		} else {
			i--
			j = 999
		}
	}
}