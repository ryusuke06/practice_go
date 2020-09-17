packege main

import (
	"fmt"
	"strconv"
)

var k int = 1 // パッケージ、関数内での初期値として定義できる
// k := 1 ':='は関数の外では使えない

func main() {
	i := 3 // 暗黙的に型宣言できる
	c, python, java := true, false, "no!"

	fmt.Println(k, i, c, python, java) // 出力の最後に自動で改行
	fmt.Printf("%T", c) // 変数のデータ型を表示・出力の際に改行されない → 改行したいなら "%T\n" と書く

	var s string = "14" // 文字列 → 数字変換
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v", i, i)

	j := make([]int, 0, 5) // make(型, len, cap)で宣言可 capは省略可(自動でlenと合わせる)
	fmt.Printf("len=%d cap=%d value=%v", len(j), cap(j), j) // len=0 cap=5 value=[0] (len(),cap()関数でlen,capをとれる)


	//連想配列

	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)									// map[apple:100 banana:200]
	fmt.Println(m["banana"])						// 200  (実は第二返り値にtrueが返ってきてるが省略可)
	m["new"] = 500
	fmt.Println(m)									// map[apple:100 banana:200 new:500]

	fmt.Println(m["nothing"])						// 0
	v, ok := m["nothing"]
	fmt.Println(v, ok)								// 0 false

	/*
	var m2 map[string]int // 初期値nil メモリー確保なし入る余地なし
	m2["pc"] = 5000
	fmt.Println(m2) // panicが起こる
	*/
	m2 := make(map[string]int) // 空のmapを作る（メモリー確保）
	m2["pc"] = 5000
	fmt.Println(m2) // map[pc:5000]


	// byte型
	a := []byte("HI") // [72 73]   []byte{72 73}でも良い
	fmt.Println(string(a)) // HI    アスキーコードは機械が読み取る文字列みたいなものだから直接キャストできる
}