package main

import "fmt"

func Pop(arr *[]int) {
	a := *arr
	*arr = a[:len(a)-1]
}

func addNum(s *[]int, letter int) {
	*s = append(*s, letter)
}

func getRestNums(all, curr []int) []int {
	m := map[int]bool{}
	for _, a := range all {
		m[a] = true
	}
	for _, a := range curr {
		delete(m, a)
	}
	res := []int{}
	for key, _ := range m {
		res = append(res, key)
	}

	return res
}

// time - O(n**n), mem - O(n*n**n)
func permute(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	var rec func(idx int, curr []int)

	rec = func(idx int, curr []int) {
		//чтобы не мутировал
		if idx == len(nums) {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			res = append(res, tmp)
			return
		}
		rest := getRestNums(nums, curr)

		for _, num := range rest {
			addNum(&curr, num)
			rec(idx+1, curr)
			Pop(&curr)

		}
	}
	rec(0, []int{})

	return res
}

func main() {
	n := []int{5, 4, 6, 2}
	fmt.Println(permute(n))
}
