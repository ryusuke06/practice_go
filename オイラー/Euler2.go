package main

func main() {
	var n int
	a := []int {1, 2}

	for i = 0; i =< 10; i++ {
		fib := append(a[i],a[i+1],a[i]+a[i+1])
		if fib[i]%2 == 0 {
			n += fib[i]
		}
	}
	println(n)
}