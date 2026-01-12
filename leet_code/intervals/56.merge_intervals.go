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
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}
	for _, curr := range intervals[1:] {
		last := merged[len(merged)-1]
		a1, b1, a2, b2 := curr[0], curr[1], last[0], last[1]

		if max(a1, a2) <= min(b1, b2) {
			merged[len(merged)-1][0], merged[len(merged)-1][1] = min(a1, a2), max(b1, b2)
		} else {
			merged = append(merged, curr)
		}

	}
	return merged
}

func main() {
	intervals := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	fmt.Println(merge(intervals))
}
