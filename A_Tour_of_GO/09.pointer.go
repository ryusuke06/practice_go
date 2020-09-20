package main

import "fmt"

func one(x *int){ // ここのasterisk(*)はポインタのint型という意味
	*x = 1 // ここの*はポインタのint型が指す値の意味
}

func main() {
	var m *int
	fmt.Println(m) // nil(初期値) なぜか　→ 宣言しただけでmemoryを確保していないため（確保したいならnew(int)で）

	var n int = 100
	one(&n) // ampersand(&)のnはnのポインタを表す
	// 1 ポインタを取らずにnを弄っても呼び出した関数内でのことだがポインタで値を指して変更すると呼び出した関数内の値も変わる
	fmt.Println(n)
}