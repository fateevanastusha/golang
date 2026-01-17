package main

import "fmt"

// time - O(n), mem - O(n)
/*
	поиск цикла.
	если попали в серую - значит цикл есть.
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	/*
		1 - white
		2 - gray
		3 - black
	*/
	matrix := [][]int{}
	for i := 0; i < numCourses; i++ {
		matrix = append(matrix, []int{})
	}
	for _, v := range prerequisites {
		a, b := v[0], v[1]
		matrix[a] = append(matrix[a], b)

	}

	colors := []int{}
	for i := 0; i < numCourses; i++ {
		colors = append(colors, 0)
	}

	var hasCycle func(n int) bool
	hasCycle = func(n int) bool {
		if colors[n] == 1 {
			return true
		}
		if colors[n] == 2 {
			return false
		}
		colors[n] = 1
		for _, next := range matrix[n] {
			if hasCycle(next) {
				return true
			}
		}
		colors[n] = 2

		return false
	}

	for n := 0; n < numCourses; n++ {
		if hasCycle(n) {
			return false
		}
	}

	return true
}

func main() {
	n := 4
	c := [][]int{{1, 2}, {2, 3}, {4, 3}}
	fmt.Println(canFinish(n, c))
}
