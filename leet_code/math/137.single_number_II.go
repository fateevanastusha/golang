package main

import (
	"fmt"
	"math"
)

// time - O(n), mem - O(1)
func singleNumber(nums []int) int {
	bytes := make([]int, 64)
	gap := int(math.Pow(2, 31))
	for _, n := range nums {
		//делаем плюс большой гап чтобы всегда работать только с положительными числами!
		//если пришло отрицательное оно будет левее 2**31, положительное справа
		n += gap
		i := 0
		for n != 0 {
			bytes[i] += n % 2
			n /= 2
			i++
		}
	}
	var r int
	for i, v := range bytes {
		r += v % 3 * int(math.Pow(2, float64(i)))
	}

	return r - gap
}

func main() {
	n := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(n))
}
