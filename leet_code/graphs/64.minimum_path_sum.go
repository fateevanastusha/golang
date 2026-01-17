package main

import (
	"fmt"
	"math"
)

// time - O(n*m), mem - O(n*m + n + m)
// если можно менять входящие данные - то память будет O(1)
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	//делаем копию грида
	dp := grid
	row := make([]int, n+1)
	//добавляем доп.ряд вверх и доп. колонку слева, состоящие из maxInt.
	for i := range row {
		row[i] = math.MaxInt
	}
	dp = append([][]int{row}, dp...)
	for i := range dp {
		dp[i] = append([]int{math.MaxInt}, dp[i]...)
	}

	//обходим со сдвигом вправо вниз
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			}
			//получаем минимальную сверху и слева, складываем с минимальной
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + dp[i][j]
		}
	}
	/*
		1 3 1
		1 5 1
		4 2 1

		1 4 1
		1 5 1
		4 2 1

		1 4 5
		1 5 1
		4 2 1

		1 4 5
		2 5 1
		4 2 1

		1 4 5
		2 7 1
		4 2 1

		1 4 5
		2 7 6
		4 2 1

		1 4 5
		2 7 6
		6 2 1

		1 4 5
		2 7 6
		6 8 1

		1 4 5
		2 7 6
		6 8 7

		получается, что каждая клетка преобразуется в сумму этой клетки и минимального числа слева и сверху.
		ответ лежит в правом нижнем углу.

	*/

	return dp[m][n]
}

func main() {
	m := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(m))
}
