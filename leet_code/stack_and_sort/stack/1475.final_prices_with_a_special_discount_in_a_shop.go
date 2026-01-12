package main

import "fmt"

func Pop(arr *[][]int) []int {
	a := *arr
	element := a[len(a)-1]
	*arr = a[:len(a)-1]
	return element
}

// time - O(n), mem - O(n)
func finalPrices(prices []int) []int {
	stack := [][]int{[]int{0, prices[0]}}

	i := 1
	for len(stack) != 0 {
		if i >= len(prices) {
			v := Pop(&stack)
			prices[v[0]] = v[1]
		} else {
			curr := prices[i]
			last := stack[len(stack)-1]
			if curr <= last[1] {
				Pop(&stack)
				prices[last[0]] = last[1] - curr
			}
			if len(stack) == 0 || curr > last[1] {
				stack = append(stack, []int{i, curr})
				i++
			}
		}

	}

	return prices
}

func main() {
	p := []int{8, 4, 6, 2, 3}
	fmt.Println(finalPrices(p))
}
