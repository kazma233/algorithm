package list

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func logLinkedList(tag string, t *testing.T, head *ListNode) {
	t.Logf("%s: %v", tag, head.Val)
	next := head.Next
	for next != nil {
		t.Logf("%s: %v", tag, next.Val)
		next = next.Next
	}
}
