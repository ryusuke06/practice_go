package main

import "fmt"

func main() {
	var a,b int = 1, 2
	var n int = 0

	for b < 4000000 {
		if b%2 == 0 {
			n += b
		}
		a,b = b, a+b
	}
	fmt.Println(n)
}
