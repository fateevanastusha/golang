package main

import (
	"fmt"
)

func good(arr []int, i int) bool {
	if i == 0 {
		return true
	}
	return arr[i-1] < arr[i]
}

// time - O(logn), mem - O(1)
func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for r-l > 1 {
		m := (l + r) / 2
		if good(arr, m) {
			l = m
		} else {
			r = m
		}
	}
	return l
}

func main() {

	a := []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}

	fmt.Println(peakIndexInMountainArray(a))

}
