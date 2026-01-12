package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// time - O(n), mem - O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummyNode := &ListNode{
		Val:  0,
		Next: head,
	}
	var l int
	curr := dummyNode
	for curr != nil {
		l++
		curr = curr.Next
	}

	curr = dummyNode
	for i := 0; i < l-n-1; i++ {
		curr = curr.Next
	}

	curr.Next = curr.Next.Next

	return dummyNode.Next
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println("nil")
}

func main() {

	// h := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 			Next: &ListNode{
	// 				Val: 4,
	// 				Next: &ListNode{
	// 					Val: 5,
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	n := 2
	res := removeNthFromEnd(h, n)
	printList(res)
}
