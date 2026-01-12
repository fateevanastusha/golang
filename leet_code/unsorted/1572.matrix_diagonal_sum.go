package main

import "fmt"

func diagonalSum(mat [][]int) int {

	sum := 0
	l := len(mat)
	for i := 0; i < len(mat); i++ {
		leftCol := i
		rightCol := i
		leftRow := i
		rightRow := l - i - 1
		if leftRow != rightRow {
			sum += mat[leftCol][leftRow] + mat[rightCol][rightRow]
		} else {
			sum += mat[leftCol][leftRow]
		}
	}

	return sum
}

func main() {

	// mat := [][]int{
	// 	{1, 2, 3, 4, 5},
	// 	{4, 5, 6, 3, 2},
	// 	{7, 8, 9, 5, 9},
	// 	{1, 3, 9, 4, 9},
	// 	{2, 6, 9, 8, 9},
	// }

	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// mat := [][]int{
	// 	{1, 1, 1, 1},
	// 	{1, 1, 1, 1},
	// 	{1, 1, 1, 1},
	// 	{1, 1, 1, 1},
	// }

	fmt.Println(diagonalSum(mat))
}
