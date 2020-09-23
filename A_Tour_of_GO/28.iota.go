package main

import "fmt"

const (
	c1 = iota
	c2 // 2番目以降のiotaは省略可能(本来はc2 = iota)
	c3
)

const (
	_ = iota
	KB int = 1 << (10 * iota) // KB = 1024 = 1 << 10 * 1
	MB // MB = 1024*1024 = 1 << 10 * 2 まるっと省略可能
	GB
)

func main()  {
	fmt.Println(c1,c2,c3) // 0 1 2  c1からの連番になる
	fmt.Println(KB, MB, GB) // 1024 1048576 1073741824
}