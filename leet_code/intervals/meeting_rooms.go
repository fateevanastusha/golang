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

// time: O(n * log n) - из-за сортировки O(n log n) + O(n) = O(n log n), mem: O(n)
func solve(intervals [][]int) int {
	/*
		Метод точек:
			[1.......................18]
				[2.........11]
					[4............15]
						[5...13]
								    [15.....................29]
									    [18........23]
			получаем так:

			1 +
			2 +
			4 +
			5 +
			11 -
			13 -
			15 -
			15 +
			18 -
			18 +
			23 -
			29 -

			- = 6
			+ = 6

			и количество одновременных (максимальное количество идущих подряд +):
			1 2 3 4 3 2 1 2 1 2 1 0
			= максимальное - 4
	*/
	points := [][]int{}
	for _, v := range intervals {
		points = append(points, []int{v[0], 1})
		points = append(points, []int{v[1], -1})
	}

	sort.Slice(points, func(i int, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	curr, max := 0, 0
	for _, point := range points {
		curr += point[1]
		if curr > max {
			max = curr
		}
	}

	return max
}

func main() {
	r := [][]int{{1, 18},
		{18, 23},
		{15, 29},
		{4, 15},
		{2, 11},
		{5, 13},
	}
	fmt.Println(solve(r))
}
