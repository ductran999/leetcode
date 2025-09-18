package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) add(node ListNode) {
	l.Next = &node
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}
	return prev
}

func main() {
	head := &ListNode{Val: 2}
	head.add(ListNode{Val: 3})

	rev := reverseList(head)

	for rev != nil {
		fmt.Print(rev.Val)
		if rev.Next != nil {
			fmt.Print("->")
		}
		rev = rev.Next
	}
}
