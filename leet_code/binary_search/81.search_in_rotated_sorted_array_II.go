package main

import "fmt"

func good(arr []int, middle int) bool {

	return
}
func search(nums []int, target int) bool {
	//find offset
	l, r := -1, len(nums)-1
	for r-l > 1 {
		m := (l + r) / 2
		if good(nums, m) {
			l = m
		} else {
			r = m
		}
	}

	// offset = r
	return false

	//find target
}

func main() {
	n := []int{2, 5, 6, 0, 0, 1, 2}
	t := 0
	fmt.Println(search(n, t))
}
