package main

import "fmt"

// 0 ms
func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	res := []int{}
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		v := nums2[i]
		if m[v] > 0 {
			res = append(res, v)
			m[v] = -1
		}
	}
	return res
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	fmt.Println(intersection(nums1, nums2))
}
