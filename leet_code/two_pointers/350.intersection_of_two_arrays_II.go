package main

import (
	"fmt"
	"slices"
)

// func intersect(nums1 []int, nums2 []int) []int {
// 	res := []int{}

// 	seen := map[int]int{}

// 	for _, v := range nums1 {
// 		seen[v]++

// 	}
// 	for _, v := range nums2 {
// 		c, ok := seen[v]
// 		if ok && c > 0 {
// 			res = append(res, v)
// 			seen[v]--
// 		}
// 	}

// 	return res
// }

func intersect(nums1 []int, nums2 []int) []int {
	slices.Sort(nums1)
	slices.Sort(nums2)
	res := []int{}

	l1, l2 := 0, 0

	for l1 < len(nums1) && l2 < len(nums2) {
		if nums1[l1] < nums2[l2] {
			l1++
		} else if nums1[l1] > nums2[l2] {
			l2++
		} else {
			res = append(res, nums1[l1])
			l1++
			l2++
		}
	}

	return res
}

func main() {

	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	fmt.Println(intersect(nums1, nums2))

}
