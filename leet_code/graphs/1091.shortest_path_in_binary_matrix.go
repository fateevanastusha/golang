package main

import (
	"fmt"
)

//time - O(n*m), mem - O(n*m)
//bfs
/*
	начинаем с 0,0 клетки. идем либо вправо/влево/вниз/вверх/вправо-вверх/вправо-вниз/влево-вверх/влево-вниз, при этом храним число
	шагов, которое необходимо, чтобы дойти до этой клетки (храним в очереди).
	bfs реализуем ОЧЕРЕДЬЮ. bfs так как он расползается равномерно во всех направлениях, это нам и нужно чтобы построить маршрут.
*/
func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}

	//make used matrix
	used := make([][]bool, len(grid))
	i := 0
	for i < len(grid) {
		used[i] = make([]bool, len(grid[0]))
		i++
	}

	//steps for bfs
	steps := [][]int{
		{0, 1},   //right
		{0, -1},  //left
		{-1, 0},  //up
		{1, 0},   //down
		{1, 1},   //down right
		{-1, -1}, //up left
		{-1, 1},  //up right
		{1, -1},  //down left
	}
	queue := [][]int{[]int{1, 0, 0}}
	used[0][0] = true

	for len(queue) != 0 {
		step, m, n := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]
		if m == len(grid)-1 && n == len(grid[0])-1 {
			return step
		}
		for _, s := range steps {
			m, n := m+s[0], n+s[1]
			//клетка в границах графа + клетка не посещена + клетка является сушей
			//если что-то из этого не так - скипаем эту клетку
			if (m >= 0 && m <= len(grid)-1 && n >= 0 && n <= len(grid[0])-1 && used[m][n] == false && grid[m][n] == 0) == false {
				continue
			}
			queue = append(queue, []int{step + 1, m, n})
			used[m][n] = true
		}
	}
	return -1
}

func main() {
	grid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	fmt.Println(shortestPathBinaryMatrix(grid))
}
