package main

import "fmt"

func do(i interface{})  {
	/*
	ii := i.(int) // Type Assertionする（interface型を他の型に変換すること）≠ cast
	ii *= 2 // ii = i * 2と一緒
	fmt.Println(ii)
	*/

	switch v := i.(type) { // switchとtypeはセット → どこかにtypeを定義するわけではない
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T", v)
	}
}

func main(){
	do(10)
	do("Mike")
	do(true)
	// interface型はどんな型も受け入れる = なのでdo("Mike")やdo(true)でもOK!

	var i int = 10
	ii := float64(10) // castもしくはtypeConversion
	fmt.Println(i, ii)
}