package main

import "fmt"

func Pop[T comparable](arr *[]T) T {
	a := *arr
	element := a[len(a)-1]
	*arr = a[:len(a)-1]
	return element
}

// time - O(n), mem - O(1)
func dailyTemperatures(temperatures []int) []int {

	/*
		стэк - это индексы температур, для которой еще не нашлось бОльшей температуры.
		мы проходимся по значениям. если стэк не пуст - смотрим, больше ли текущее значение
		последнего значения в стэке, снимаем его. (повторяем пока стэк не опустеет либо пока
		последнее значение в нем не станет >= текущего).
	*/
	res := make([]int, len(temperatures))
	stack := []int{0}
	i := 1
	for len(stack) != 0 {
		//значения закончились, остались те, что без пары. им ставим 0
		if i >= len(temperatures) {
			stackLast := Pop(&stack)
			res[stackLast] = 0
		} else {
			//текущее значение
			curr := temperatures[i]
			//последнее из стэка
			stackLast := stack[len(stack)-1]
			//если текущее значение больше последнего из стэка
			if curr > temperatures[stackLast] {
				//снимаем со стека, пишем найденную пару
				Pop(&stack)
				res[stackLast] = i - stackLast
			}
			//либо мы опустошили стэк, либо последнее значение не меньше текущего. кладем текущее.
			if len(stack) == 0 || curr <= temperatures[stackLast] {
				stack = append(stack, i)
				i++
			}
		}

	}
	return res
}

func main() {
	//[1,1,4,2,1,1,0,0]

	//8
	t := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(t))
}
