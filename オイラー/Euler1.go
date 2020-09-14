package main

func main() {
	var n int
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			n += i
		}
	}
	println(n)
}

//A.233168