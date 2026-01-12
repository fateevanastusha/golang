package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// time: O(n), mem: 0(n)
func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	// s - mid
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	swappedMid := &ListNode{
		Val:  s.Val,
		Next: nil,
	}

	curr := s.Next
	for curr != nil {
		swappedMid = &ListNode{
			Val:  curr.Val,
			Next: swappedMid,
		}
		curr = curr.Next
	}

	a, b := head, swappedMid
	for b != nil {
		if b.Val != a.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}

	//лучше вернуть в исходный вид

	return true

}

func main() {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	fmt.Println(isPalindrome(h))
}
