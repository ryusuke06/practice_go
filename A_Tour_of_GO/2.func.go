package main

import "fmt"

func add(x, y int) int {
	return x + y
}
// 返り値にINT型の引数を足したINT型の値を返す関数として定義

func swap(x, y string) (string, string) {
	return y, x
}
// 複数の返り値が戻ってくる場合(データ型, データ型)で定義

func split(sum int) (x, y int) { // 引数に名前(ここではsum)と名付けることができる
	x = sum * 4 / 9
	y = sum - x
	return // 引数の変数を使うとreturnステートメントに何も書かなくて良い（naked return）
}

func main() {
	fmt.Println(add(42, 13)) // 55

	a, b := swap("hello", "world")
	fmt.Println(a, b) // world hello

	fmt.Println(split(17))
}
