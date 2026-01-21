package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func clear(lists []*ListNode) []*ListNode {

	j := 0
	for _, v := range lists {
		if v != nil {
			lists[j] = v
			j++
		}
	}
	return lists[:j]
}

func find(lists []*ListNode) int {
	//ищет индекс наименьшей
	minValue := math.MaxInt
	k := 0
	for i, v := range lists {
		if v != nil && v.Val < minValue {
			minValue = v.Val
			k = i
		}
	}

	return k
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for {
		//на каждрй итерации цикла находим наименьшую ноду и добавляем ее
		//удаляет nil-значения
		lists = clear(lists)
		if len(lists) == 0 {
			break
		}
		//находим индекс ноды с наименьшим значением
		point := find(lists)

		tail.Next = lists[point]
		lists[point] = lists[point].Next
		tail = tail.Next
	}

	return dummy.Next

}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println("nil")
}

func main() {

	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}

	res := mergeKLists([]*ListNode{l1, l2, l3})
	printList(res)

}
