packege main

import "fmt"

var k int = 1 // パッケージ、関数内での初期値として定義できる
// k := 1 ':='は関数の外では使えない

func main() {
	i := 3 // 暗黙的に型宣言できる
	c, python, java := true, false, "no!"

	fmt.Println(k, i, c, python, java)
}