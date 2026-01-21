package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// time: O(n), mem: 0(n)
func isPalindrome(head *ListNode) bool {

	/*
		1 > 2 > 2 > 1 - true
		1 > 2 > 3 > 4 > 4 > 3 > 2 > 1 - true


	*/
	if head == nil || head.Next == nil {
		return true
	}

	//сначала через fast/slow pointers находим середину
	//1 > 2 > 3 > 4 > 4 > 3 > 2 > 1
	//4 > 3 > 2 > 1 - середина
	// s - mid
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	//dummy node для середины
	swappedMid := &ListNode{
		Val:  s.Val,
		Next: nil,
	}

	//разворачиваем вторую половину
	//4 > 3 > 2 > 1
	//1 > 2 > 3 > 4

	curr := s.Next
	for curr != nil {
		swappedMid = &ListNode{
			Val:  curr.Val,
			Next: swappedMid,
		}
		curr = curr.Next
	}

	//два указателя, один с начала, другой с РАЗВЕРНУТОЙ середины
	//если ноды не равны - значит не палиндром
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
