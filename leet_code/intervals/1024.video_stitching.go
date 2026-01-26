package main

import (
	"fmt"
	"sort"
)

// time - O(nlogn), mem - O(n)
func videoStitching(clips [][]int, time int) int {

	//сортируем по началу, если они равны, то по бОльшему концу (чтобы он был первее)
	sort.Slice(clips, func(a, b int) bool {
		aa, bb := clips[a], clips[b]
		if aa[0] == bb[0] {
			return aa[1] > bb[1]
		}
		return aa[0] < bb[0]
	})

	//самый маленький должен быть 0, иначе эта область никогда не покроется
	if clips[0][0] != 0 {
		return -1
	}
	curr := 1               //количество клипов
	currTime := clips[0][1] //время которое надо покрыть
	lastClip := 0           //индекс последнего клипа

	//жадный алгоритм
	for currTime < time {
		// start, end := clips[lastClip][0], clips[lastClip][1]

		//для текущего клипа ищем лучший для перехода
		best := -1
		for i := lastClip + 1; i < len(clips); i++ {
			//пропускаем те что уже не входят в интервал
			nextClip := clips[i]
			if nextClip[1] <= currTime {
				continue
			}
			//начало следующего уже вышло за интервал, выходим, дальше идти смысла нет (так как дальше начало будет больше)
			//вышли за границу времени
			if nextClip[0] > currTime {
				break
			}
			//он больше последнего клипа
			if nextClip[1] > clips[lastClip][1] {
				//проверяем что он лучше текущего best, если нет - скипаем
				if best != -1 && nextClip[1] < clips[best][1] {
					continue
				}
				//если лучше - назначаем его лучшим
				best = i
			}
		}
		if best == -1 {
			return -1
		}

		//расширяем текущее время
		currTime = clips[best][1]
		//добавляем счетчик ответа
		curr++
		//последний клип - лучший
		lastClip = best
	}

	//не покрыли до конца
	if currTime < time {
		return -1
	}

	return curr
}

func main() {
	c := [][]int{
		{8, 10}, {17, 39}, {18, 19}, {8, 16}, {13, 35}, {33, 39}, {11, 19}, {18, 35},
	}
	t := 20
	fmt.Println(videoStitching(c, t))
}
