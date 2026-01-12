package main

import "fmt"

func setZeroes(matrix [][]int) {

	m, n := len(matrix), len(matrix[0])
	zeroes := [][]int{}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == 0 {
				zeroes = append(zeroes, []int{row, col})
			}
		}
	}

	nulledRows, nulledCols := map[int]bool{}, map[int]bool{}

	for _, v := range zeroes {
		row, col := v[0], v[1]

		if !nulledRows[row] {
			//zero row

			for i, _ := range matrix[row] {
				matrix[row][i] = 0
			}
			nulledRows[row] = true

		}
		if !nulledCols[col] {
			//zero col
			for i, _ := range matrix {
				matrix[i][col] = 0
			}

			nulledCols[col] = true
		}

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
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	setZeroes(matrix)

	printMatrix(matrix)

}
