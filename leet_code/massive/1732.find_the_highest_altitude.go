package main

import (
	"fmt"
)

// time - O(n), mem - O(n)
func largestAltitude(gain []int) int {
	prefix := append([]int{0}, gain...)
	res := 0

	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i] + prefix[i-1]
		res = max(prefix[i], res)
	}

	return res
}

func main() {

	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}
