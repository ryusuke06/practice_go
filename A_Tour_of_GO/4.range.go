package main

import "fmt"

func main() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++{
		fmt.Println(i, l[i]) // 0 python 1 go 2 java
	}

	for i, v := range l{
		fmt.Println(i, v) // 0 python 1 go 2 java
	} // indexが必要ないならアンダースコア

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m{
		fmt.Println(k, v) // apple 100 banana 200
	} // keyが必要ない時はアンダースコア　※　valueが必要ない時はそのまま省略でOK
}