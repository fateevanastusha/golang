package main

import "fmt"

// time - 0(logn), mem - O(1)
/*
	здесь мы по сути матрицу:
	1 2 3
	4 5 6
	7 8 9

	воспринимаем как обычный массив:
	1 2 3 4 5 6 7 8 9


	у нас l = 0, r = 9, ответ в левом указателе
*/
func searchMatrix(matrix [][]int, target int) bool {

	l, r := 0, len(matrix)*len(matrix[0])
	m := len(matrix[0])
	for r-l > 1 {
		mid := (l + r) / 2
		/*
			1 2 3
			4 5 6
			7 8 9
			например, mid = 3 (это число 4)
			row = 3 / 3 = 1
			col = 3 % 3 = 0

			все верно
		*/
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
