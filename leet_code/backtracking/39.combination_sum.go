package main

import (
	"fmt"
	"sort"
)

func Pop(arr *[]int) {
	a := *arr
	*arr = a[:len(a)-1]
}

func Add(s *[]int, l int) {
	*s = append(*s, l)
}

// time - O(n**n)
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var rec func(start int, curr []int, currSum int)
	sort.Ints(candidates)

	rec = func(start int, curr []int, currSum int) {
		if currSum > target {
			return
		}
		if currSum == target {
			temp := make([]int, len(curr))
			copy(temp, curr)
			res = append(res, temp)
			return
		}

		/*
			каждый раз начинаем не с начала, а с того же числа (так как числа могут повторяться)
			благодаря этому не будет повторений комбинаций
		*/
		for i := start; i < len(candidates); i++ {
			num := candidates[i]
			Add(&curr, num)
			rec(i, curr, currSum+num)
			Pop(&curr)

		}
	}
	rec(0, []int{}, 0)

	return res
}

func main() {
	c := []int{2, 3, 6, 7}
	t := 7
	fmt.Println(combinationSum(c, t))
}
