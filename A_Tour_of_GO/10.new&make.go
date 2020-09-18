package main

import "fmt"


func main() {
	s := make([]int, 0)
	fmt.Printf("%T\n", s) // []int

	m := make(map[string]int)
	fmt.Printf("%T\n", m) // map[string]int

	ch := make(chan int)
	fmt.Printf("%T\n", ch) // chan int

	var p *int = new(int) // memory確保 ちなみに*pでポインタが指す値は0
	fmt.Printf("%T\n", p) // *int ポインタを返す値にはnew()関数をつかう

	var st = new(struct{})
	fmt.Printf("%T\n", st) // *struct{}
}