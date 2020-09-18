package main

import "fmt"

type Vertex struct {
	x,y int // 小文字でx, yの型リテラルはこのパッケージ内のみ（x,yが大文字で他のパッケージからも見えるようになる）
}

func Area(v Vertex)int{ // 関数
	return v.x * v.y
}

func (v Vertex) Area() int{ // メソッドで結びつきのある値 → 値レシーバー
	// メソッド 名前の前に結びつきのある型をparenthesesで示す(Area()からわかるように引数ではない → レシーバー)
	return v.x * v.y
}

func New(x, y int) *Vertex{
	// 他のパッケージから生成したこのパッケージを読み込むときにstructを型リテラルが読み取れない中、値を返すために必要な関数（コンストラクタ）GOのよくあるパターン
	return &Vertex{x, y}
}

func (v *Vertex) Scale(i int) { // メソッドで結びつきのある値がポインタ → ポインタレシーバー
	/*
	引数にiをとり、レシーバーは*Vertex（Vertexの値はドット演算子で繋いだ値、勝手にポインタに直して計算！）
	asteriskを外せば呼び出した関数の中身を弄らずポインタから値を受け取ってることに変える！！
	*/
	v.X = v.x * i
	v.Y = v.y * i
}

func main() {
	// v := Vertex{3, 4}
	v := New(3, 4)
	fmt.Println(Area(v)) // 関数はvの結びつきがない（使用するときに値を渡す必要がある）
	v.Scale(10) // 引数は10、*Vertexはvから勝手に取って来る！！！！しゅごい！
	fmt.Println(v.Area()) // メソッドは結びつきがある（ドット演算子で呼び出せるのでわかりやすい）
}