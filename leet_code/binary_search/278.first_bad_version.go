package main

import "fmt"

func isBadVersion(version int) bool {
	return version >= 1
}

// 0ms
func firstBadVersion(n int) int {

	low := 1
	high := n

	for low <= high {
		mid := (low + high) / 2
		if isBadVersion(mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}

func main() {

	fmt.Println(firstBadVersion(3))

}

//[0,1,2,3,4,5]
//[0,1,2,3,4,5,6,7,8,9,10]
