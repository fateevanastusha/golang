package main

import (
	"fmt"
	"math"
)

// в этой задаче может быть две точки в одном месте! если так - то в map храним массив [y, cnt], а в оптимизации понижаем cnt!
// time - 0(n), mem - O(n)
func lineReflection(points [][]int) bool {
	maxX, minX := math.MinInt, math.MaxInt
	p := map[int]int{}

	for _, point := range points {
		x := point[0]
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		p[point[0]] = point[1]
	}

	//check if every has pair
	for _, point := range points {
		x, y := point[0], point[1]
		//if no in hash table - we already check his pair, skip
		if _, ok := p[x]; !ok {
			continue
		}

		//ФОРМУЛА ЧТОБЫ НАЙТИ ПРОТИВОПОЛОЖНЫЙ Х!!! МОЖНО ВЫВЕСТИ НА ГРАФИКЕ НА БУМАЖКЕ
		xd := maxX + minX - x

		if yd, ok := p[xd]; !ok || (ok && y != yd) {
			return false
		} else {
			//delete from hash table, to not check again his pair
			delete(p, xd)
			delete(p, x)
		}
	}

	return true
}
func main() {
	p := [][]int{{1, 1}, {2, 1}, {4, 1}, {5, 1}}
	fmt.Println(lineReflection(p))
}
