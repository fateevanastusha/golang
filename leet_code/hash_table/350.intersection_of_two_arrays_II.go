package main

import (
	"fmt"
)

// time - O(n), mem - O(n)
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		return intersect(nums2, nums1)
	}
	var res []int
	m := map[int]int{}
	for _, v := range nums1 {
		m[v]++
	}

	for _, v := range nums2 {
		if m[v] > 0 {
			res = append(res, v)
			m[v]--
		}
	}

	return res
}

func main() {

	nums1 := []int{9, 4, 9, 8, 4}
	nums2 := []int{4, 9, 5}

	fmt.Println(intersect(nums1, nums2))

}
