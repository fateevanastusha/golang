package main

import (
	"fmt"
)

// time - O(n) (n - длина входного числа), mem - O(1)
func reverse(x int) int {
	if x < 0 {
		return -reverse(-x)
	}
	if x >= 2147483647 {
		return 0
	}

	var res int
	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}

	if res >= 2147483647 {
		return 0
	}

	return res
}

func main() {
	n := 12345

	fmt.Println(reverse(n))
}
