package main

import (
	"fmt"
	"sort"
)

func isIncl(otr []int, point int) bool {
	return point >= otr[0] && point <= otr[1]
}

// time - O(n*m*logn), mem - O(n*m)
func solve(otr [][]int, points []int) []int {

	p := [][]int{}

	for _, o := range otr {
		p = append(p, []int{o[0], +1})
		p = append(p, []int{o[1], -1})
	}

	for i, o := range points {
		p = append(p, []int{o, 0, i})
	}

	sort.Slice(p, func(i int, j int) bool {

		if p[i][0] == p[j][0] {
			if p[i][1] == 0 {
				return false
			}
			return p[i][1] < p[j][1]
		}
		return p[i][0] < p[j][0]
	})

	currSum, currN := 0, 0

	for _, v := range p {
		if v[1] == 0 {
			points[v[2]] = currSum
			currN++
		}
		currSum += v[1]
	}

	return points
}

func main() {
	otr := [][]int{{-10, 10}, {3, 5}}
	points := []int{1, 2, 3}
	fmt.Println(solve(otr, points))

}
