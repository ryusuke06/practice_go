package main

import "fmt"

/*
もっともよく使われているinterfaceの一つに fmt パッケージ に定義されている Stringer があります:

type Stringer interface {
    String() string
}

Stringer インタフェースは、stringとして表現することができる型です。
fmt パッケージ(と、多くのパッケージ)では、変数を文字列で出力するためにこのインタフェースがあることを確認します。
*/

type Person struct{
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %v.", p.Name) // My name is Mike.
	// stringのメソッドを実装すればAgeを返さない文字列を作れたりもする
}

func main(){
	var mike = Person{"Mike", 22}
	fmt.Println(mike) // {Mike 22} これを独自に書き換えたいときにstringerを使う
}