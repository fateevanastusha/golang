package main

import (
	"fmt"
	"slices"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	values := []int{}
	curr := head
	for curr != nil {
		values = append(values, curr.Val)
		curr = curr.Next
	}

	slices.Sort(values)
	h := &ListNode{}
	tail := h
	for i := 0; i < len(values); i++ {
		tail.Val = values[i]
		if i != len(values)-1 {
			tail.Next = &ListNode{}
			tail = tail.Next
		}

	}

	return h
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println("nil")
}

func main() {

	l := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}

	res := sortList(l)
	printList(res)
}
