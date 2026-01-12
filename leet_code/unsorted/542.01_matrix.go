package main

import "fmt"

func updateMatrix(mat [][]int) [][]int {
	var findNearest func(row int, col int, who string) int
	findNearest = func(row int, col int, who string) int {
		v := mat[row][col]
		if v == 0 {
			return 0
		}
		const INF = int(1e9)
		m := INF
		if col > 0 && who != "right" {
			d := 1 + findNearest(row, col-1, "left")
			if d < m {
				m = d
			}
		}
		if col < len(mat[0])-1 && who != "left" {
			d := 1 + findNearest(row, col+1, "right")
			if d < m {
				m = d
			}
		}
		if row > 0 && who != "down" {
			d := 1 + findNearest(row-1, col, "up")
			if d < m {
				m = d
			}
		}
		if row < len(mat)-1 && who != "up" {
			d := 1 + findNearest(row+1, col, "down")
			if d < m {
				m = d
			}

		}
		return m
	}
	for row := 0; row < len(mat); row++ {
		for column := 0; column < len(mat[row]); column++ {
			if mat[row][column] != 0 {
				fmt.Println(row, column)
				mat[row][column] = findNearest(row, column, "")
			}
		}
	}

	return mat

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

	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}

	printMatrix(updateMatrix(mat))
}
