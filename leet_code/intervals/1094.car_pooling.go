package main

import (
	"fmt"
	"sort"
)

func carPooling(trips [][]int, capacity int) bool {
	var max, curr int
	var points [][]int
	for _, v := range trips {
		points = append(points, []int{v[1], v[0]})
		points = append(points, []int{v[2], -v[0]})
	}

	sort.Slice(points, func(i int, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	for _, point := range points {
		curr += point[1]
		if curr > max {
			max = curr
		}
	}

	return max <= capacity
}

func main() {
	//[4,5,6],[6,4,7],[4,3,5],[2,3,5]
	trips := [][]int{{4, 5, 6}, {6, 4, 7}, {4, 3, 5}, {2, 3, 5}}
	capacity := 13
	fmt.Println(carPooling(trips, capacity))
}

/*
          [5....6]
     [4..............7]
  [3....5]
  [3....5]

  + 2
  + 4
  + 6
  - 2
  - 4
  + 4
  - 4
  - 6

*/
