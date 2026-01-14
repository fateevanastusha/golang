package main

import (
	"fmt"
)

// time - O(n), mem - O(max)
func hIndex(citations []int) int {
	maxCit := 0
	for _, v := range citations {
		if v > maxCit {
			maxCit = v
		}
	}
	cnt := make([]int, maxCit+1)
	for _, v := range citations {
		cnt[v]++
	}

	papers := 0
	for h := maxCit; h >= 0; h-- {
		papers += cnt[h]
		if papers >= h {
			return h
		}
	}
	return 0
}

func main() {
	n := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(n))
}
