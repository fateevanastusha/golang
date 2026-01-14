package main

import "fmt"

// time - O(n*m), mem - O(n*m)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	//if stone in start position
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[1][1] = 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//если это камень - пропускаем. на месте камней в dp будет 0
			if (i == 1 && j == 1) || obstacleGrid[i-1][j-1] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}

func main() {
	n := [][]int{{1}}
	fmt.Println(uniquePathsWithObstacles(n))
}
