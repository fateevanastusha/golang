package main

import (
	"fmt"
	"math"
)

func getValue(arr []int, i int) int {
	if i < 0 {
		return math.MinInt
	}
	return arr[i]
}

// time - O(m+n), mem - O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	for p1 >= 0 || p2 >= 0 {
		i := p1 + p2 + 1
		v1, v2 := getValue(nums1, p1), getValue(nums2, p2)
		if v1 >= v2 {
			nums1[i] = v1
			p1--
		} else {
			nums1[i] = v2
			p2--
		}
	}
}

func main() {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	m := 3
	nums2 := []int{3, 2, 1}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

//{1,2,0,0,0}
//m=2
//{1,2,3}
//n=3
