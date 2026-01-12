package main

import "fmt"

func rotate(matrix [][]int) {
	values := [][]int{}

	for r, row := range matrix {
		for c, v := range row {
			values = append(values, []int{r, c, v})
		}
	}

	n := len(matrix) - 1

	for _, c := range values {
		r, c, v := c[0], c[1], c[2]
		newRow := c
		newCol := n - r

		matrix[newRow][newCol] = v
	}

}

func printMatrix(mat [][]int) {
	for _, row := range mat {
		for _, v := range row {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotate(matrix)

	printMatrix(matrix)

	fmt.Println()

	matrix2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}

	rotate(matrix2)

	printMatrix(matrix2)

}
