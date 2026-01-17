package main

import "fmt"

/*
	Что лучше?
	dfs писать проще
	bfs еще на очередь тратит память
	чем плох dfs?
		1) stack вызовов - обычно в районе 1000 (его может не хватить и все что больше 10*100 вызывает вопросы)
		2)
*/

// time - O(n), mem - O(m)
/*
	если можем менять входящий массив - то можем не использовать массив used, а просто красить клетку в воду
	то есть делать вместо:
	used[m][n] = true
	использовать:
	grid[m][n] = "0"



	КАК ЭТО ВСЕ РЕШАЕТСЯ:
	делаем матрицу с посещенными клетками, при посещении каждой клетки отмечаем ее как true, и больше мы в нее не попадаем.
	итерируемся по всем клеткам, для каждой (если она есть остров и не посещена) - вызываем dfs, он проходится по острову,
	каждый раз после прохода прибавляем.
	по проходу dfs:
	1) вниз
	2) влево
	3) вверх
	4) вправо
	если клетка за пределами grid, посещена или является водой - пропускаем ее.
*/

//DFS

// func numIslands(grid [][]byte) int {
// 	//make used matrix
// 	used := make([][]bool, len(grid))
// 	i := 0
// 	for i < len(grid) {
// 		used[i] = make([]bool, len(grid[0]))
// 		i++
// 	}

// 	//dfs func
// 	var dfs func(m, n int)
// 	dfs = func(m, n int) {
// 		//клетка в границах графа + клетка не посещена + клетка является сушей
// 		//если что-то из этого не так - скипаем эту клетку
// 		if (m >= 0 && m <= len(grid)-1 && n >= 0 && n <= len(grid[0])-1 && used[m][n] == false && grid[m][n] == '1') == false {
// 			return
// 		}
// 		used[m][n] = true
// 		steps := [][]int{
// 			//down
// 			{m + 1, n},
// 			//left
// 			{m, n - 1},
// 			//up
// 			{m - 1, n},
// 			//right
// 			{m, n + 1},
// 		}
// 		for _, s := range steps {
// 			dfs(s[0], s[1])

// 		}
// 	}

// 	var cnt int
// 	for m := range grid {
// 		for n := range grid[m] {
// 			if used[m][n] || grid[m][n] == '0' {
// 				continue
// 			}
// 			dfs(m, n)
// 			cnt++
// 		}
// 	}
// 	return cnt
// }

//DFS with queue
// func Pop(arr *[][]int) []int {
// 	a := *arr
// 	element := a[len(a)-1]
// 	*arr = a[:len(a)-1]
// 	return element
// }

// func numIslands(grid [][]byte) int {
// 	//make used matrix
// 	used := make([][]bool, len(grid))
// 	i := 0
// 	for i < len(grid) {
// 		used[i] = make([]bool, len(grid[0]))
// 		i++
// 	}

// 	//dfs func
// 	var dfs func(m, n int)
// 	dfs = func(m, n int) {
// 		stack := [][]int{[]int{m, n}}
// 		used[m][n] = true
// 		for len(stack) != 0 {
// 			last := Pop(&stack)
// 			m, n := last[0], last[1]
// 			steps := [][]int{
// 				//down
// 				{m + 1, n},
// 				//left
// 				{m, n - 1},
// 				//up
// 				{m - 1, n},
// 				//right
// 				{m, n + 1},
// 			}
// 			for _, s := range steps {
// 				//клетка в границах графа + клетка не посещена + клетка является сушей
// 				//если что-то из этого не так - скипаем эту клетку
// 				if (s[0] >= 0 && s[0] <= len(grid)-1 && s[1] >= 0 && s[1] <= len(grid[0])-1 && used[s[0]][s[1]] == false && grid[s[0]][s[1]] == '1') == false {
// 					continue
// 				}
// 				stack = append(stack, s)
// 				used[s[0]][s[1]] = true
// 			}
// 		}

// 	}

// 	var cnt int
// 	for m := range grid {
// 		for n := range grid[m] {
// 			if used[m][n] || grid[m][n] == '0' {
// 				continue
// 			}
// 			fmt.Println(m, n)
// 			bfs(m, n)
// 			cnt++
// 		}
// 	}
// 	return cnt
// }

//BFS with queue
// func numIslands(grid [][]byte) int {
// 	//make used matrix
// 	used := make([][]bool, len(grid))
// 	i := 0
// 	for i < len(grid) {
// 		used[i] = make([]bool, len(grid[0]))
// 		i++
// 	}

// 	//bfs func
// 	var bfs func(m, n int)
// 	bfs = func(m, n int) {
// 		queue := [][]int{[]int{m, n}}
// 		used[m][n] = true
// 		for len(queue) != 0 {
// 			last := queue[0]
//			queue = queue[1:]
// 			m, n := last[0], last[1]
// 			steps := [][]int{
// 				//down
// 				{m + 1, n},
// 				//left
// 				{m, n - 1},
// 				//up
// 				{m - 1, n},
// 				//right
// 				{m, n + 1},
// 			}
// 			for _, s := range steps {
// 				//клетка в границах графа + клетка не посещена + клетка является сушей
// 				//если что-то из этого не так - скипаем эту клетку
// 				if (s[0] >= 0 && s[0] <= len(grid)-1 && s[1] >= 0 && s[1] <= len(grid[0])-1 && used[s[0]][s[1]] == false && grid[s[0]][s[1]] == '1') == false {
// 					continue
// 				}
// 				queue = append(queue, s)
// 				used[s[0]][s[1]] = true
// 			}
// 		}

// 	}

// 	var cnt int
// 	for m := range grid {
// 		for n := range grid[m] {
// 			if used[m][n] || grid[m][n] == '0' {
// 				continue
// 			}
// 			fmt.Println(m, n)
// 			bfs(m, n)
// 			cnt++
// 		}
// 	}
// 	return cnt
// }

func main() {
	//1 - land, 0 - water
	grid := [][]byte{
		{'0', '0', '1', '0', '0'},
		{'0', '1', '1', '0', '0'},
		{'0', '1', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
		{'1', '0', '0', '0', '0'},
	}

	fmt.Println(numIslands(grid))
}
