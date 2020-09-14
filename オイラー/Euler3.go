package main

import "fmt"

func main() {
	var a int = 600851475143
	for i := 2; i < 100000; {
    		if a%i == 0 {
    			a = a/i
				if a == 1 {
					fmt.Println(i)
					break
				}
    		}else if a%i != 0 {
    			i++
		}
	}
}
