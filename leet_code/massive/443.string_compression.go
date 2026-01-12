package main

import (
	"fmt"
	"strconv"
)

// time: O(n), mem: O(1)
func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}
	writePointer, currLength := 0, 0
	temp := append(chars, '1')

	for readPointer := 1; readPointer <= len(chars); readPointer++ {
		curr := chars[readPointer-1]
		currLength++
		if readPointer == len(chars) || curr != temp[readPointer] {
			chars[writePointer] = curr
			writePointer++
			if currLength == 1 {
				currLength = 0
				continue
			}
			for _, v := range []byte(strconv.Itoa(currLength)) {
				chars[writePointer] = v
				writePointer++
			}
			currLength = 0
		}
	}

	return writePointer
}

func main() {
	// test1 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	// fmt.Println(compress(test1), string(test1))

	// test2 := []byte{'a'}
	// fmt.Println(compress(test2), string(test2))

	// test3 := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	// fmt.Println(compress(test3), string(test3))

	test4 := []byte{'a', 'a', 'a', 'a', 'a', 'b'}
	fmt.Println(compress(test4), string(test4))
}
