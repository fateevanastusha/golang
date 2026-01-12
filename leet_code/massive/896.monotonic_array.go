package main

import "fmt"

func isMonotonic(nums []int) bool {

	//time: O(n), mem: O(1)
	asc := true
	desc := true

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			desc = false
		}
		if nums[i] < nums[i-1] {
			asc = false
		}
	}

	return asc || desc

	// if len(nums) == 1 || len(nums) == 2 {
	// 	return true
	// }
	// isStarted := false
	// isDown := false
	// i := 0
	// for i < len(nums)-1 {
	// 	curr, next := nums[i], nums[i+1]

	// 	if curr > next {
	// 		if !isStarted {
	// 			isDown = true
	// 			isStarted = true
	// 		} else if isDown == false {
	// 			return false
	// 		}
	// 	}
	// 	if curr < next {

	// 		if !isStarted {
	// 			isStarted = true
	// 		} else if isDown == true {
	// 			return false
	// 		}

	// 	}
	// 	i++
	// }

	// return true
}

func main() {
	n := []int{1, 2, 4, 2}
	fmt.Println(isMonotonic(n))
}
