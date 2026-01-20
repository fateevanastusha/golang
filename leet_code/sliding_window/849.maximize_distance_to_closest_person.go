package main

import "fmt"

// time - O(n), mem - O(1)
func maxDistToClosest(seats []int) int {
	l, r := 0, 0
	res := 0
	for l < len(seats) {

		for r+1 < len(seats) && seats[r] == seats[r+1] {
			r++
		}
		if seats[r] == 0 {

			var d int
			if l == 0 || r == len(seats)-1 {
				res = max(r-l+1, res)
			} else {
				res = max((r-l+2)/2, res)
			}

			res = max(d, res)

		}

		l = r + 1
		r = r + 1
	}

	return res
}

func main() {
	n := []int{1, 0, 0, 0}
	fmt.Println(maxDistToClosest(n))
}
