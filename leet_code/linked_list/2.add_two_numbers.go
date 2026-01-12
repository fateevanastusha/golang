package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// time: O(n), mem: O(n)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	carry := 0
	curr := dummy

	for l1 != nil || l2 != nil || carry != 0 {
		var l1Val, l2Val int

		if l1 != nil {
			l1Val = l1.Val
		}
		if l2 != nil {
			l2Val = l2.Val
		}

		sum := l1Val + l2Val + carry
		newNode := &ListNode{
			Val: sum % 10,
		}
		carry = sum / 10
		curr.Next = newNode
		curr = curr.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

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

	l1 := &ListNode{Val: 2, Next: &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	}}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	res := addTwoNumbers(l1, l2)
	printList(res)
}
