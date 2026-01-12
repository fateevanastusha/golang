package main

import "fmt"

func searchMatrixII(matrix [][]int, target int) bool {

	m, n := len(matrix)-1, len(matrix[0])-1
	row, col := 0, n

	for row <= m && col >= 0 {
		curr := matrix[row][col]
		if curr == target {
			return true
		}
		if curr > target {
			col--
		}
		if curr < target {
			row++
		}
	}

	return false
}

func main() {

	/*
		+----+----+----+----+----+
		|  1 |  4 |  7 | 11 | 15 |
		+----+----+----+----+----+
		|  2 |  5 |  8 | 12 | 19 |
		+----+----+----+----+----+
		|  3 |  6 |  9 | 16 | 22 |
		+----+----+----+----+----+
		| 10 | 13 | 14 | 17 | 24 |
		+----+----+----+----+----+
		| 18 | 21 | 23 | 26 | 30 |
		+----+----+----+----+----+
	*/

	m := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	t := -10

	fmt.Println(searchMatrixII(m, t))
}
