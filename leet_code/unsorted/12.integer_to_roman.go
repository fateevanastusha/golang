package main

import "fmt"

func intToRoman(num int) string {

	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	letters := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""

	for i := 0; i < len(nums); i++ {
		for num >= nums[i] {
			res += letters[i]
			num -= nums[i]
		}
	}

	return res
}

func main() {

	n := 1994

	fmt.Println(intToRoman(n))

}
