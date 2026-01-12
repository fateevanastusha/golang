package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	var prevSlow *ListNode
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prevSlow = slow
		slow = slow.Next
	}

	return prevSlow
}

func reverseList(head *ListNode) *ListNode {

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

func reorderList(head *ListNode) {
	if head.Next == nil {
		return
	}
	mid := middleNode(head)
	p2 := reverseList(mid.Next)
	mid.Next = nil
	p1 := head

	for p2 != nil {
		p1Next := p1.Next
		p1.Next = p2
		p1 = p2
		p2 = p1Next
	}

}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println("nil")
}

func main() {
	n := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			// 	Next: &ListNode{
			// 		Val: 3,
			// 		Next: &ListNode{
			// 			Val: 4,
			// 			Next: &ListNode{
			// 				Val: 5,
			// 				Next: &ListNode{
			// 					Val: 6,
			// 				},
			// 			},
			// 		},
			// 	},
		},
	}

	reorderList(n)
	printList(n)

}
