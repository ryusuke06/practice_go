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

func split(x, y  int) (sum int) { // 返り値に名前(ここではsum)と名付けることができる
	sum = x + y // sumはすでに定義されているので":="ではなく"="
	return // 返り値を名付けるとreturnステートメントに何も書かなくて良い（naked return）
}

// クロージャ
func later() func(string) string {
	// １つ前に与えられた文字列を保存するための変数（関数に関係する関数がいの環境）
	var store string
	// 引数に文字列を取り文字列を返す関数を返す(クロージャ)
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// 可変長引数
func foo(params ...int){ // 可変長は"..."で表す
	fmt.Println(len(params), params) // foo(10,20) だと 2 [10, 20]
	for _ , param := range params{
		fmt.Println(param) // 10 20
	}
}

func main() {
	fmt.Println(add(42, 13)) // 55

	a, b := swap("hello", "world")
	fmt.Println(a, b) // world hello

	fmt.Println(split(17))

	f := func(x, y int){
		fmt.Println(x,y)
	}(1,2) // 関数の中に関数を呼ぶ → 関数も同じデータ型、特別扱いしない(外で定義してmainで呼び出すのと一緒じゃん)

	// クロージャ（関数と関数外の環境をセットにして閉じ込めたもの）
	fmt.Println(later("Golang")) // ""
	fmt.Println(later("is")) // "Golang"

	// 可変長引数
	s := []int{1, 2, 3}
	foo(s...)  // スライスの中身を1,2,3と分解して渡すこともできる
}
