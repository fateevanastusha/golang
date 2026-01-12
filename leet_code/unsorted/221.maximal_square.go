package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix[0]), len(matrix)

	res := 0

	fmt.Println(n)

	for row := 0; row < n-1; row++ {
		for col := 0; col < m-1; col++ {
			leftT, leftB, rightT,rightB := matrix[row][col], matrix[]
		}
	}

	return res
}

func main() {

	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix))
}
