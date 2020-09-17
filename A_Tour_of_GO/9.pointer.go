package main

import "fmt"

func one(x *int){ // ここのasterisk(*)はポインタのint型という意味
	*x = 1 // ここの*はポインタのint型が指す値の意味
}

func main() {
	var n int = 100
	one(&n) // ampersand(&)のnはnのポインタを表す
	// 1 ポインタを取らずにnを弄っても呼び出した関数内でのことだがポインタで値を指して変更すると呼び出した関数内の値も変わる
	fmt.Println(n)
}