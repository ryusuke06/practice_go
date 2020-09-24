/*
13195 の素因数は 5, 7, 13, 29 である.

600851475143 の素因数のうち最大のものを求めよ.
*/

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

//A.6857