package main

import "fmt"

// time: 0(n), mem: O(1)
func fib(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b

	// if n == 1 || n == 0 {
	// 	return n
	// }
	// return fib(n-1) + fib(n-2)
}

func main() {
	n := 3
	fmt.Println(fib(n))
}
