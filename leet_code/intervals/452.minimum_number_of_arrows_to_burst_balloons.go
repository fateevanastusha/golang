package main

import (
	"fmt"
	"sort"
)

func ifIntersect(a1, b1, a2, b2 int) bool {
	return max(a1, a2) <= min(b1, b2)
}

// time: O(n * log(n)), mem: O(1)
/*
	тут нужно понять сколько надо стрел чтобы проткнуть все интервалы минимально



                        [10───────────────────────16]
      [2────────8]
  [1──────6]
              [7────────────────────12]

	получается, что тут 2:
	1-6 + 2-8
	10-16 + 7-12
*/

func findMinArrowShots(points [][]int) int {

	//сортируем интервалы по началу
	sort.Slice(points, func(i int, j int) bool {
		return points[i][0] < points[j][0]
	})

	//количество стрел (минимум 1)
	curr := 1
	//текущий интервал
	interval := points[0]

	for _, point := range points[1:] {

		a1, b1, a2, b2 := point[0], point[1], interval[0], interval[1]

		//если текущий пересекается с последним, то назначаем их ПЕРЕСЕЧЕНИЕ (не объединение)
		//тогда 1-6 и 2-8 превратятся в 2-6
		//тогда мы всегда будем знать, что стрела проткнет самый левый точно (то есть всех из этого объединения)
		if ifIntersect(a1, b1, a2, b2) {
			interval[0] = max(a2, a1)
			interval[1] = min(b1, b2)
		} else {
			//если не пересекаются - вытаскиваем новую стрелу для следующей группы интервалов
			curr++
			interval = point
		}
	}

	return curr
}

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(points))
}
