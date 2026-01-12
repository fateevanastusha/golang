package main

import "fmt"

func convert(s string, numRows int) string {

	rs := []rune(s)
	l := len(rs)
	if l <= 2 || numRows == 1 {
		return s
	}
	res := make([]rune, 0, l)
	for row := 0; row < numRows; row++ {
		swap := false
		getGap := func() int {
			gap := (numRows - 2) + numRows
			if row == 0 || row == numRows-1 {
				return gap //up and down
			}
			if swap {
				swap = false
				return 2 * row //near
			} else {
				swap = true
				return gap - (2 * row) //distant
			}
		}
		for i := row; i < l; i += getGap() {
			res = append(res, rs[i])
		}
	}

	return string(res)
}

func main() {
	s := "PAYPALISHIRING"
	r := 4
	fmt.Println(convert(s, r))
}
