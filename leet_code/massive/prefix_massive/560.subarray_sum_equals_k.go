package main

import "fmt"

// ????
// time: 0(n), mem: O(n)
func solve(nums []int, k int) int {
	var cnt, currPx int
	px := map[int]int{0: 1}

	for _, v := range nums {
		currPx += v
		cnt += px[currPx-k]
		px[currPx]++
	}

	return cnt
}

func main() {

	nums := []int{0, 0, 1, 2, 3}
	k := 3

	fmt.Println(solve(nums, k))

}
