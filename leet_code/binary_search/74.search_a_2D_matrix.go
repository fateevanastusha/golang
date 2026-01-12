package main

import "fmt"

// time - 0(logn), mem - O(1)
func searchMatrix(matrix [][]int, target int) bool {

	l, r := 0, len(matrix)*len(matrix[0])
	m := len(matrix[0])
	for r-l > 1 {
		mid := (l + r) / 2
		row, col := mid/m, mid%m
		if matrix[row][col] <= target {
			l = mid
		} else {
			r = mid
		}
	}

	return matrix[l/m][l-(m*(l/m))] == target
}

func main() {

	m := [][]int{{1}, {3}, {5}, {7}}
	t := 5

	fmt.Println(searchMatrix(m, t))
}
