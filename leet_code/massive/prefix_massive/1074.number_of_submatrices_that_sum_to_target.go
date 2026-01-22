package main

import "fmt"

// time - O(m*m*n), mem - O(m*n)
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	prefixMatrix := matrix
	//добавляем ряд сверху и колонку слева
	for i := range prefixMatrix {
		prefixMatrix[i] = append([]int{0}, prefixMatrix[i]...)
	}
	row := make([]int, len(prefixMatrix[0]))
	prefixMatrix = append([][]int{row}, prefixMatrix...)

	//индексируем матрицу
	for a := 1; a < len(prefixMatrix); a++ {
		for b := 1; b < len(prefixMatrix[0]); b++ {
			prefixMatrix[a][b] = prefixMatrix[a-1][b] + prefixMatrix[a][b-1] + prefixMatrix[a][b] - prefixMatrix[a-1][b-1]
		}

	}

	var cnt int

	//проходим все комбинации row
	for r1 := 1; r1 < len(prefixMatrix); r1++ {
		for r2 := r1; r2 < len(prefixMatrix); r2++ {
			//считаем префиксную сумму для стоблца
			/*
					0 0 0 0
					0 1 2 3
				r1	0 2 4 6
				r2	0 3 6 9  -> 3-1=2 6-2=4 9-3=6


						0 0 0 0
				r2 r1 	0 0 1 1
						0 1 3 4
						0 1 4 5  -> 0 1 1
			*/
			px := map[int]int{0: 1}
			for c := 1; c < len(prefixMatrix[0]); c++ {
				//берем текущую префиксную сумму (она уже посчитана, просто вычитаем из нее всех, что выше)
				currPx := prefixMatrix[r2][c] - prefixMatrix[r1-1][c]
				cnt += px[currPx-target]
				px[currPx]++
			}
		}

	}

	return cnt
}

func main() {
	m := [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}
	t := 0
	fmt.Println(numSubmatrixSumTarget(m, t))
}
