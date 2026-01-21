package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// time - O(n), mem - 0(1)
/*
	После обнаружения цикла (встречи fast и slow), сбрасываем один указатель в начало списка,
	а второй оставляем в точке встречи. Двигаем оба по шагу за раз. Место их новой встречи —
	начало цикла.
*/
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow

}

func main() {

	node := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}

	// node.Next.Next.Next.Next = node.Next
	fmt.Println(detectCycle(node))
}
