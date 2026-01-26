package main

import "fmt"

// time - O(logn), mem - O(1)
func search(nums []int, target int) int {

	l, r := 0, len(nums)

	/*
		ищем насколько произошел сдвиг
		//4, 5, 6, 7, | 0, 1, 2, 3
			GOOD			BAD

		GOOD все больше или равны первому элементу
		BAD все меньше первого элемента
	*/
	//find pivot
	for r-l > 1 {
		m := (l + r) / 2
		if nums[m] > nums[0] {
			l = m
		} else {
			r = m
		}
	}

	offset := l + 1
	l, r = 0, len(nums)

	for r-l > 1 {
		m := (l + r) / 2
		//получаем индекс так, как он должен был быть если бы не сместили массив
		/*
			4, 5, 6, 7, | 0, 1, 2, 3

			offset = 4

			если middle 7
			сейчас у него индекс - 3
			должен быть - 7

			3 + 4 % 8 = 7 % 8 = 7


			если middle 2
			сейчас у него индекс - 6
			должен быть - 2
			6 + 4 % 8 = 10 % 8 = 2


			и по сути обычный бинарный поиск

			в крайнем случае можно сделать из rotated массива нормальный.
		*/
		i := (m + offset) % len(nums)
		if nums[i] <= target {
			l = m
		} else {
			r = m
		}
	}

	result := (l + offset) % len(nums)

	if nums[result] != target {
		return -1
	}
	return result
}

func main() {
	m := []int{4, 5, 6, 7, 0, 1, 2, 3}
	t := 1
	fmt.Println(search(m, t))
}
