package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func ifIntersect(a1, b1, a2, b2 int) bool {
	return max(a1, a2) <= min(b1, b2)
}

// time: O(n * log(n)), mem: O(1)
func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i int, j int) bool {
		return points[i][0] < points[j][0]
	})

	curr := 1
	interval := points[0]

	for _, point := range points[1:] {

		a1, b1, a2, b2 := point[0], point[1], interval[0], interval[1]

		if ifIntersect(a1, b1, a2, b2) {
			interval[0] = max(a2, a1)
			interval[1] = min(b1, b2)
		} else {
			curr++
			interval = point
		}
	}

	return curr
}

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(points))
}
