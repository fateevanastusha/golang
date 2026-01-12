// package main

// import "fmt"

// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	write := 1
// 	for read := 1; read < len(nums); read++ {
// 		if nums[read] != nums[read-1] {
// 			nums[write] = nums[read]
// 			write++
// 		}
// 	}
// 	return write
// }

// func main() {
// 	arr := []int{1, 1, 2}
// 	fmt.Println(removeDuplicates(arr))
// }

package main

import "fmt"

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	write := 1
	last := nums[0]
	for r := 1; r <= len(nums)-1; r++ {
		value := nums[r]
		if value != last {
			nums[write] = nums[r]
			write++
		}
		last = nums[r]

	}

	return write
}

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr))
	fmt.Println(arr)
}
