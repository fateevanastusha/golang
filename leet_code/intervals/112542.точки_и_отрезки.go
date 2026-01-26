package main

import (
	"fmt"
	"sort"
)

func isIncl(otr []int, point int) bool {
	return point >= otr[0] && point <= otr[1]
}

/*
	Дается на вход:
	интервалы и точки

	вернуть нужно для каждой точки в какое количество отрезков она входит
*/
// time - O(n*m*logn), mem - O(n*m)
func solve(otr [][]int, points []int) []int {

	p := [][]int{}

	//добавляем точки с отрезков
	for _, o := range otr {
		p = append(p, []int{o[0], +1})
		p = append(p, []int{o[1], -1})
	}

	//добавляем точки с точек
	for i, o := range points {
		p = append(p, []int{o, 0, i}) //[сама_точка, обозначение_что_это_точка, индекс_точки]
	}

	// Сортируем события по начальной координате.
	// Если начальные координаты равны, порядок такой:
	// 1) конец отрезка (-1)
	// 2) точка (0)
	// 3) начало отрезка (+1)
	//
	// Это нужно для отрезков вида [l, r]:
	// – при x = r точка должна считаться внутри отрезка
	// – при x = l точка должна считаться внутри отрезка
	sort.Slice(p, func(i int, j int) bool {

		if p[i][0] == p[j][0] {
			if p[i][1] == 0 {
				return false
			}
			return p[i][1] < p[j][1]
		}
		return p[i][0] < p[j][0]
	})

	//в каждый момент знаем сколько отрезков сейчас активны
	currSum := 0

	for _, v := range p {
		//если точка, то мы выяснили в какое кол-во отрезков она входит
		if v[1] == 0 {
			//по индексу ставим количество "активных отрезков" в данный момент этой точке
			points[v[2]] = currSum
		}
		currSum += v[1]
	}

	return points
}

func main() {
	otr := [][]int{{-10, 10}, {3, 5}}
	points := []int{1, 2, 3}
	fmt.Println(solve(otr, points))

}
