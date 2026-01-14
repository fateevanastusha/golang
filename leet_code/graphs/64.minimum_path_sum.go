package main

import (
	"fmt"
	"math"
)

// time - O(n*m), mem - O(n*m + n + m)
// если можно менять входящие данные - то память будет O(1)
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := grid
	row := make([]int, n+1)
	for i := range row {
		row[i] = math.MaxInt
	}
	dp = append([][]int{row}, dp...)
	for i := range dp {
		dp[i] = append([]int{math.MaxInt}, dp[i]...)
	}

	dp[1][1] = grid[0][0]
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 {
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + dp[i][j]
		}
	}
	return dp[m][n]
}

func main() {
	m := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(m))
}
