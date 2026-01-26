package main

import (
	"fmt"
	"math"
)

// в этой задаче может быть две точки в одном месте! если так - то в map храним массив [y, cnt], а в оптимизации понижаем cnt!
/*
	ПО СУТИ НАМ НАДО РАЗДЕЛИТЬ X ЛИНИЕЙ ОТНОСИТЕЛЬНО КОТОРОЙ ТОЧКИ СПРАВА И СЛЕВА ЛЕЖАТ ОДИНАКОВО.
	для каждой точки нам по сути надо найти противоположный x (где должен располагаться его пара), и проверить, что их y - равны.
	как найти противоположный х?

		y
		^
		4|				 |
		3|       *  x1   |   x2  *
		2|           *   |   *
		1|				 |
		0+---|---|---|---|---|---|---> x
			1   2   3   4   5   6

	 	l (середина) - это ((max-min)/2) + min = 4
		значит, чтобы найти x1:

		x1=l-(x2-l)
		=2l-x2
		=2(((max-min)/2) + min)-x2
		=max-min+2min-x2
		=max+min-x2

		x2=(l-x1)+l
		=2l-x1
		...(то же самое что выше)
		=max+min-x1

		то есть чтобы найти противоположную точку - max-min-x2

*/
// time - 0(n), mem - O(n)
func lineReflection(points [][]int) bool {
	maxX, minX := math.MinInt, math.MaxInt
	//key - x, value - y
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

	//проверяем каждое ли число на той стороне имеет пару. чекаем каждую точку один раз, не имеет смысла оба значения из пары проверять.
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
