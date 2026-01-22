package main

import "fmt"

func matrixBlockSum(mat [][]int, k int) [][]int {
	var res [][]int
	res = append(res, mat...)

	//add zeroes
	/*
		добавлю нули слева и сверху (ряд и колонку)
	*/
	for i, v := range mat {
		mat[i] = append([]int{0}, v...)
	}
	z := []int{}
	for _, _ = range mat[0] {
		z = append(z, 0)
	}
	mat = append([][]int{z}, mat...)

	//индексирую массив
	/*
		0 0 0 0
		0 1 1 1
		0 1 1 1
		0 1 1 1

		0 0 0 0
		0 1 2 3
		0 2 4 6
		0 3 6 9

		(получается что каждый квадрат это сумма всех чисел)

		0 0 0 | 0
		0 1 2 | 3
		0 2 4 | 6
		------
		0 3 6   9

		складываем верхнее и левое число, и вычитаем по диагонали (так как мы его два раза добавили) + добавляем
		собственное число этой клетки
		4 = 2 + 2 - 1 + 1
		2 + 2 - левое и верхнее
		- 1 - по диагонали
		+ 1 - собственное число клетки
	*/
	for row := 1; row < len(mat); row++ {
		for col := 1; col < len(mat[0]); col++ {
			mat[row][col] = mat[row][col] + mat[row-1][col] + mat[row][col-1] - mat[row-1][col-1]
		}
	}

	n := len(mat) - 1
	m := len(mat[0]) - 1

	/*
		теперь для каждой клетки выделяем квадрат, вычетаем из него лишние квадраты и добавляем по диагонали
		(так как вычли два раза). это и будет результат
		0 0   0 0
		0 1   2 3
		     ----
		0 2 | 4 6
		0 3 | 6 9

		например, чтобы найти для нижнего правого квадрата 2x2 сумму элементов -
		вычитаем левый (3), вычитаем верхний (3), добавляем диагональный (1)
		9 - 3 - 3 + 1 = 4
	*/
	for row := 1; row < len(mat); row++ {
		for col := 1; col < len(mat[0]); col++ {
			//leftUp
			row1, col1 := row-k, col-k
			//rightDown
			row2, col2 := row+k, col+k
			//border cases
			if row1 < 1 {
				row1 = 1
			}
			if row2 > n {
				row2 = n
			}
			if col1 < 1 {
				col1 = 1
			}
			if col2 > m {
				col2 = m
			}

			//find sum of this square
			sum := mat[row2][col2] - mat[row1-1][col2] - mat[row2][col1-1] + mat[row1-1][col1-1]
			res[row-1][col-1] = sum
		}
	}

	return res
}

func printSquare(mat [][]int, row1, row2, col1, col2 int) {
	fmt.Println("Square:")
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			fmt.Printf("%4d", mat[i][j])
		}
		fmt.Println()
	}
	fmt.Println("------")
}

func main() {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	k := 1
	fmt.Println(matrixBlockSum(m, k))
}
