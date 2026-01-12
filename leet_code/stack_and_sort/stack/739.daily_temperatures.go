package main

import "fmt"

func Pop[T comparable](arr *[]T) T {
	a := *arr
	element := a[len(a)-1]
	*arr = a[:len(a)-1]
	return element
}

func dailyTemperatures(temperatures []int) []int {

	res := make([]int, len(temperatures))
	stack := []int{0}
	i := 1
	for len(stack) != 0 {
		if i >= len(temperatures) {
			stackLast := Pop(&stack)
			res[stackLast] = 0
		} else {
			//get curr value
			curr := temperatures[i]
			//get last of stack
			stackLast := stack[len(stack)-1]
			if curr > temperatures[stackLast] {
				//снимаем со стека, пишем найденную пару
				Pop(&stack)
				res[stackLast] = i - stackLast
			}
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
