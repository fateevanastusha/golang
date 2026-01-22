package main

import "fmt"

/*
найти повторяющееся и пропущенное число в последовательности [1,n]
совершенно логичное уравнение:
OS - original sum (какая должна быть)
CS - current sum (какая сейчас)
x - пропущенное
y - повторяющееся
двойное уравнение:
OS = CS + x - y (плюс пропущенное минус повторяющееся)
OS^2= CS^2 + x^2 - y^2

OS - CS = x - y
OS^2 - CS^2 = x^2 - y^2

A = OS - CS
B = OS^2 - CS^2

A = x - y
B = x^2 - y^2

для y:
x = A + y
B = (A + y)^2 - y^2 >
B = A^2 + 2*A*y + y^2 - y^2 >
B = A^2 + 2*A*y >
2*A*y = B - A^2 >
...
y = (B - A^2)/2*A

для x:
y = x - A
B = x^2 - (x - A)^2
B = x^2 - (x^2 - 2*A*x + A^2)
B = x^2 - x^2 + 2*A*x - A^2
B = 2*A*x - A^2
B + A^2 = 2*A*x
...
x = (B + A^2)/(2*A)


дальше:

*/
//time - O(n), mem - O(1)
func repeatedNumber(m []int) []int {
	//SQ_OS = OS^2, SQ_CS = CS^2
	OS, CS, SQ_OS, SQ_CS := 0, 0, 0, 0
	for i, v := range m {
		CS += v
		SQ_CS += v * v
		SQ_OS += (i + 1) * (i + 1)
		OS += i + 1
	}
	A := CS - OS
	B := SQ_CS - SQ_OS
	x, y := (B+A*A)/2*A, (B-A*A)/(2*A)
	return []int{y, x}
}

func main() {
	ints := []int{3, 1, 2, 5, 3}
	fmt.Println(repeatedNumber(ints))
}
