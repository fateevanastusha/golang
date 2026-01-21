package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// time: 0(n), mem: 0(1)
func middleNode(head *ListNode) *ListNode {

	/*
		есть быстрый и медленный указатель - быстрый шагает через 2 шага, медленный - через 1
		↓
		1 > 2 > 3 > 4 > 5 > 6 > nil
		↑

		    ↓
		1 > 2 > 3 > 4 > 5 > 6 > nil
		        ↑

		        ↓
		1 > 2 > 3 > 4 > 5 > 6 > nil
		                ↑

		            ↓
		1 > 2 > 3 > 4 > 5 > 6 > nil
		                         ↑

		как только быстрый достигает конца, медленный как раз на середине.
	*/
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func main() {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}

	fmt.Println(middleNode(h))

}
