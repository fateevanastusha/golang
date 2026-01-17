package main

import (
	"fmt"
)

// time - O(n), mem - O(n)
/*
	топологическая сортировка.
	тут так же есть проверка цикла.
	при окрашивании вершины в черную - пушим ее в список в конец. таким образом порядок получится верный.
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	/*
		топологическая сортировка
	*/
	matrix := make([][]int, numCourses)
	colors := make([]int, numCourses)
	for _, n := range prerequisites {
		a, b := n[0], n[1]
		matrix[a] = append(matrix[a], b)
	}

	res := []int{}

	var dfs func(n int) bool
	dfs = func(n int) bool {
		if colors[n] == 1 {
			return true
		}
		if colors[n] == 2 {
			return false
		}
		colors[n] = 1
		for _, c := range matrix[n] {
			//check if we have cycle
			if dfs(c) {
				return true
			}
		}
		colors[n] = 2
		res = append(res, n)
		return false
	}

	for c := 0; c < numCourses; c++ {
		if dfs(c) {
			return []int{}
		}
	}
	return res
}

func main() {
	n := 4
	c := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	fmt.Println(findOrder(n, c))
}
