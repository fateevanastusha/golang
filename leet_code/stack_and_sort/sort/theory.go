package main

import (
	"fmt"
	"math"
)

//BUBBLE COUNT:

/*
1 проход:
[9 8 7 6 5 4]
[8 9 7 6 5 4]
[8 7 9 6 5 4]
[8 7 6 9 5 4]
[8 7 6 5 9 4]
[8 7 6 5 4 9]
САМОЕ БОЛЬШОЕ ЧИСЛО ПЕРЕШЛО В КОНЕЦ

2 проход:
[8 7 6 5 4 9]
[7 8 6 5 4 9]
[7 6 8 5 4 9]
[7 6 5 8 4 9]
[7 6 5 4 8 9]

3 проход:
[7 6 5 4 8 9]
[6 7 5 4 8 9]
[6 5 7 4 8 9]
[6 5 4 7 8 9]

4 проход:
[6 5 4 7 8 9]
[5 6 4 7 8 9]
[5 4 6 7 8 9]

5 проход:
[5 4 6 7 8 9]
[4 5 6 7 8 9]

сложность получилась:
n + (n-1) + (n-2)... = n^2-((1+n)/2)n=(n^2/2)-n=n^2 (округлили до n^2)
*/

// time - O(n^2), mem - O(1)
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

/*
	class Solution:
    # Bubble sort
    # time:      O(n * n)
    # mem (доп): O(1)
    def sortArray(self, nums: List[int]) -> List[int]:
        swapCount = -1
        i = 0
        while swapCount != 0:
            swapCount = 0
            for j in range(len(nums) - i - 1):
                if nums[j] > nums[j + 1]:
                    swapCount += 1
                    nums[j], nums[j + 1] = nums[j + 1], nums[j]
            i += 1
        return nums
*/
// time - O(n^2)
func bubbleSortOptimized(arr []int) []int {
	swap := -1
	i := 0

	for swap != 0 {
		swap = 0
		//так как вправо всегда приходит самый большой элемент, то не обязательно доходить до конца
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap++
			}
		}
		i++
	}

	return arr
}

/*
      [11 6 3 24   46 22 7]			ДЕЛИМ
   [11 6 3 24]       [46 22 7]		ДЕЛИМ
 [11 6]   [3 24]   [46 22]  [7]		ДЕЛИМ
[6] [11] [3] [24] [46] [22] [7]		ДЕЛИМ
 [6 11]   [3 24]   [22 46]  [7]		МЕРДЖИМ
  [3 6 11 24]        [7 22 46]		МЕРДЖИМ
      [3 6 7 11 22 24 26]			МЕРДЖИМ


	Каждый раз делим массив на два, пока он не станет размером 1
*/

func getValue(arr []int, i int) int {
	if i < 0 {
		return math.MinInt
	}
	return arr[i]
}

// time - O(n), mem - O(n)
func merge(nums1 []int, nums2 []int) []int {
	p1, p2 := len(nums1)-1, len(nums2)-1
	result := make([]int, len(nums1)+len(nums2))
	for p1 >= 0 || p2 >= 0 {
		i := p1 + p2 + 1
		v1, v2 := getValue(nums1, p1), getValue(nums2, p2)
		if v1 >= v2 {
			result[i] = v1
			p1--
		} else {
			result[i] = v2
			p2--
		}
	}
	return result
}

// time - O(n*logn), mem - O(n)
// более эффективен на большом количестве элементов
func mergeSort(nums []int) []int {
	//[8] [9] [6] [7] [4] [5] [3]
	//[8 9] [6 7] [4 5] [3]
	//[6 7 8 9] [3 4 5]
	//[3 4 5 6 7 8 9]
	if len(nums) == 1 {
		return nums
	}
	// if len(nums) <= 32 {
	// 	//оптимизация
	// 	return bubbleSortOptimized(nums)
	// }

	left := mergeSort(nums[:len(nums)/2])
	right := mergeSort(nums[len(nums)/2:])
	return merge(left, right)
}

/*
	пивотом можно выбрать либо самое левое значение, либо самое правое, либо центральное
*/

// time avg: O(n logn), worse case: O(n^2), mem avg: O(logn)
// классический
// func quickSort(nums []int) []int {
// 	m := len(nums) / 2
// 	pivot := nums[m]

// 	less, great, equal := []int{}, []int{}, []int{}
// 	for _, el := range nums {
// 		if el > pivot {
// 			great = append(great, el)
// 		} else if el < pivot {
// 			less = append(less, el)
// 		} else {
// 			equal = append(equal, el)
// 		}
// 	}

// 	result := []int{}
// 	result = append(result, quickSort(less)...)
// 	result = append(result, equal...)
// 	result = append(result, quickSort(great)...)
// 	return result
// }

// ебанутый
func getPivot(arr []int, l int, r int) int {
	//получаем рандомный пивот
	// pivotIdx := rand.IntN(r-l) + l
	//можно брать всегда крайний
	pivotIdx := r - 1
	return pivotIdx
}

func partition(nums []int, lIdx, rIdx int) int {
	pivotIdx := getPivot(nums, lIdx, rIdx)
	pivot := nums[pivotIdx]
	//ставим пивот в самое начало чтобы он не мешал сортировать
	nums[lIdx], nums[pivotIdx] = nums[pivotIdx], nums[lIdx]
	//r-1 делаем так как правый не включительно
	l, r := lIdx+1, rIdx-1
	/*
		делим на 2 партиции:
		1 - элементы <= pivot
		2 - элементы >= pivot
	*/
	for l <= r {
		if nums[l] < pivot {
			l++
		} else if nums[r] > pivot {
			r--
		} else {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	//меняем пивот с правым, чтобы справа от пивота были бОльшие числа, а слева меньшие
	nums[lIdx], nums[r] = nums[r], nums[lIdx]
	return r
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	pivot := partition(nums, l, r)
	// fmt.Println(pivot)
	quickSort(nums, l, pivot)   //не включительно правый!
	quickSort(nums, pivot+1, r) //не включительно правый!
	return
}

func main() {

	n := []int{9, 8, 7, 6, 5, 4, 3}
	fmt.Println("bubble sort -", bubbleSort(n))
	n = []int{9, 8, 7, 6, 5, 4, 3}
	fmt.Println("bubble sort optimized -", bubbleSortOptimized(n))
	n = []int{9, 8, 7, 6, 5, 4, 3}
	fmt.Println("merge sort -", mergeSort(n))
	n = []int{9, 8, 7, 5, 2, 3, 10}
	quickSort(n, 0, len(n))
	fmt.Println("quick sort -", n)

}
