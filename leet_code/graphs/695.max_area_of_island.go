package main

import "fmt"

// time - O(n*m), mem - O(n*m)
/*
	обходим все острова через dfs и сохраняем для них площадь (каждая клетка +1 к площади). затем выбираем
	максимальную площадь.
	обход через dfs:
	в порядке: вниз, влево, вверх, вправо
*/
func maxAreaOfIsland(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
	}

	var dfs func(a, b int) int
	dfs = func(a, b int) int {
		if (a >= 0 && a < len(grid) && b >= 0 && b < len(grid[0]) && visited[a][b] == false && grid[a][b] == 1) == false {
			return 0
		}
		visited[a][b] = true
		currCnt := 1
		steps := [][]int{
			//down
			{a + 1, b},
			//left
			{a, b - 1},
			//up
			{a - 1, b},
			//right
			{a, b + 1},
		}
		for _, step := range steps {
			currCnt += dfs(step[0], step[1])
		}

		return currCnt
	}

	maxCnt := 0
	for a := 0; a < len(grid); a++ {
		for b := 0; b < len(grid[a]); b++ {
			if grid[a][b] != 0 && visited[a][b] != true {
				currCnt := dfs(a, b)
				if currCnt > maxCnt {
					maxCnt = currCnt
				}
			}
		}
	}

	return maxCnt

}

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
}
