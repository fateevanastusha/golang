package main

import (
	"fmt"
)

/*
решето Эратосфена

n=10
0 1 2 3 4 5 6 7 8 9
F F T T T T T T T T

для 2:
0 1 2 3 4 5 6 7 8 9
F F T T F T F T F T

для 3:
0 1 2 3 4 5 6 7 8 9
F F T T F T F T F F

дальше скипаем и считаем сколько True

решето НЕЛЬЗЯ НАЧАТЬ С СЕРЕДИНЫ, ТОЛЬКО С ДВОЙКИ!
*/

// time - O(nloglogn), mem - O(n)
func countPrimes(n int) int {

	if n <= 1 {
		return 0
	}
	primes := make([]bool, n)
	for i := 0; i < n; i++ {
		primes[i] = true
	}
	primes[0], primes[1] = false, false

	//можно i < math.Sqrt2(n)+1
	for i := 2; i < n; i++ {
		if primes[i] == false {
			continue
		}
		for j := i * i; j < n; j += i {
			primes[j] = false
		}
	}
	var cnt int
	for _, v := range primes {
		if v {
			cnt++
		}
	}

	return cnt
}

func main() {
	n := 10
	fmt.Println(countPrimes(n))
}
