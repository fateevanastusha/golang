package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	curr := head
	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
			continue
		}
		curr = curr.Next
	}

	return head
}

func main() {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	node := (deleteDuplicates(h))
	res := []int{}

	curr := node
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}

	fmt.Println(res)
}
