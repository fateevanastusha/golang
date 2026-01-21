package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//time: O(n), mem: O(1)

func reverseList(head *ListNode) *ListNode {

	/*
		null < 1 < 2 < 3 < 4

		меняет направление
	*/
	var prev *ListNode
	curr := head
	for curr != nil {
		tmp := curr
		curr = curr.Next
		tmp.Next = prev
		prev = tmp
	}

	return prev
}

func main() {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	node := (reverseList(h))
	res := []int{}

	curr := node
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}

	fmt.Println(res)
}
