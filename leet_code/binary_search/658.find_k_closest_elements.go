package main

import "fmt"

/*
class Solution:

	# time: O(log(n - k))
	# mem:  O(k)
	def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
		# немного переформулируем задачу:
		# нас просят найти самую выгодную позицию "плавающего окна"
		# в котором K наиболее ближайших элементов к X

		# для решения будем использовать бинарный поиск
		# т к можно для каждой позиции окна однозначно сказать выгодно
		# двигать окно в право или нет

		# хороший - если подвинуть вправо выгодно
		# сравниваем arr[i] и arr[i + k] чтобы понять что ближе к "x"
		# если arr[i] ближе, то нет смысла двигать вправо т к
		# arr[i + k] элемент, который будет в нашем "плавающем окне"
		# если шагнуть вправо
		# Note: arr[i + k] - это именно элемент который войдет
		# если подвинуть плавающее окно, т е сейчас он не входит
		def good(i: int):
			return x - arr[i] > arr[i + k] - x

		# ответ будет в "r", поэтому изначально l = -1
		# чтобы r мог принимать значения [0, len(arr) - k] включительно
		l, r = -1, len(arr) - k
		while r - l > 1:
			m = (l + r) // 2
			if good(m):
				l = m
			else:
				r = m
		return arr[r:r+k]
*/
func good(arr []int, middle, target, k int) bool {
	return target-arr[middle] > arr[middle+k]-target
}
func findClosestElements(arr []int, k int, target int) []int {
	l, r := -1, len(arr)-k
	for r-l > 1 {
		m := (l + r) / 2
		if good(arr, m, target, k) {
			l = m
		} else {
			r = m
		}
	}

	return arr[r : r+k]
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 4
	target := 3
	fmt.Println(findClosestElements(nums, k, target))
}
