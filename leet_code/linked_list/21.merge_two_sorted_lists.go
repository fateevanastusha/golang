package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getVal(node *ListNode) int {
	if node == nil {
		return math.MaxInt64
	}
	return node.Val
}

// 0 ms
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	/*
		два указателя - на первый и второй список. берем наименьшее значение, пушим, двигаем указатель у того,
		что взяли.
	*/
	dummy := &ListNode{}
	tail := dummy
	for l1 != nil || l2 != nil {
		if getVal(l1) <= getVal(l2) {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
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
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l2 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	printList(mergeTwoLists(l1, l2))
}
