package main

import "fmt"

// time - O(n), mem - O(n)
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	p1, p2 := 1, 1

	for p2 < len(nums) {
		//на границе разных чисел
		if nums[p2] != nums[p2-1] {
			v := nums[p2]
			//выставляем в write указатель текущее значение
			nums[p1] = v
			//двигаем write указатель
			p1++
			//пока указатель на том же числе - двигаем его
			for p2 < len(nums) && nums[p2] == v {
				p2++
			}
		} else {
			p2++
		}
	}

	return p1
}

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr))
	fmt.Println(arr)
}
