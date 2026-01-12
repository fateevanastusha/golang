package main

import (
	"fmt"
	"sort"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// time - O(logn*m), mem - O(1)
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	cnt := 0
	sort.Ints(arr2)
	for i := 0; i < len(arr1); i++ {
		l, r := 0, len(arr2)
		f := true
		for r-l > 1 {
			m := (l + r) / 2
			if absInt(arr2[m]-arr1[i]) <= d {
				f = false
				break
			} else {
				if arr2[m] > arr1[i] {
					r = m
				} else {
					l = m
				}
			}
		}
		if absInt(arr2[l]-arr1[i]) <= d {
			f = false
		}
		if f == true {
			cnt++
		}
	}
	return cnt
}

func main() {
	arr1 := []int{3, 4, 5, 8, 9, 10, 100}
	arr2 := []int{10, 9, 1, 8}
	d := 2
	fmt.Println(findTheDistanceValue(arr1, arr2, d))
}
