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

func videoStitching(clips [][]int, time int) int {

	sort.Slice(clips, func(a, b int) bool {
		aa, bb := clips[a], clips[b]
		if aa[0] == bb[0] {
			return aa[1] > bb[1]
		}
		return aa[0] < bb[0]
	})

	min := 1
	maxLasted := clips[0][1]
	fromClip := 0

	if maxLasted >= time {
		return min
	}

	for {

	}

	// var res int

	// sort.Slice(clips, func(a, b int) bool {
	// 	aa, bb := clips[a], clips[b]
	// 	if aa[0] == bb[0] {
	// 		return aa[1] > bb[1]
	// 	}
	// 	return aa[0] < bb[0]
	// })

	// intervals := [][]int{{0, time}}
	// for _, curr := range clips {
	// 	a1, b1 := curr[0], curr[1]
	// 	forRemove := []int{}
	// 	isUsed := false
	// 	for i, interval := range intervals {
	// 		a2, b2 := interval[0], interval[1]
	// 		new := [][]int{}
	// 		if max(a1, a2) < min(b1, b2) {
	// 			if a1 > a2 && b1 < b2 {
	// 				new = append(new, []int{a1, a2})
	// 				new = append(new, []int{b1, b2})
	// 				forRemove = append(forRemove, i)
	// 			} else {
	// 				intervals[i] = []int{max(b1, a2), max(b2, a1)}
	// 			}

	// 			if intervals[i][0] == intervals[i][1] {
	// 				forRemove = append(forRemove, i)
	// 			}
	// 			isUsed = true

	// 			fmt.Println(curr)
	// 		}
	// 	}
	// 	if isUsed == true {
	// 		res++
	// 	}
	// 	for _, v := range forRemove {
	// 		intervals = append(intervals[:v], intervals[v+1:]...)
	// 	}

	// }

	// if len(intervals) != 0 {
	// 	return -1
	// }
	// return res
}

func main() {
	c := [][]int{
		{0, 1}, {6, 8}, {0, 2}, {5, 6},
		{0, 4}, {0, 3}, {6, 7}, {1, 3},
		{4, 7}, {1, 4}, {2, 5}, {2, 6},
		{3, 4}, {4, 5}, {5, 7}, {6, 9},
	}
	t := 9
	fmt.Println(videoStitching(c, t))
}
