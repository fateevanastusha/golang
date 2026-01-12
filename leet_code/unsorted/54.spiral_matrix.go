package main

import "fmt"

func idx(i, n int) int {
	if i%2 == 0 {
		return i / 2
	}
	return n - 1 - i/2
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := []int{}

	top, bottom, left, right := 0, m-1, 0, n-1

	for top <= bottom && left <= right {

		//слева направо
		for col := left; col <= right; col++ {
			res = append(res, matrix[top][col])
		}
		top++
		if top > bottom {
			break
		}

		//сверху вниз
		for row := top; row <= bottom; row++ {
			res = append(res, matrix[row][right])
		}
		right--
		if left > right {
			break
		}

		//справа налево
		for col := right; col >= left; col-- {
			res = append(res, matrix[bottom][col])
		}
		bottom--
		if top > bottom {
			break
		}

		//снизу вверх
		for row := bottom; row >= top; row-- {
			res = append(res, matrix[row][left])
		}
		left++
		if left > right {
			break
		}

	}

	return res

}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4, 5, 6},
		{7, 8, 9, 10, 11, 12},
		{13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30},
	}

	fmt.Println(spiralOrder(matrix))

}
