package list

import (
	"testing"
)

func Test_reverseLinkedList(t *testing.T) {
	three := &ListNode{
		Val:  3,
		Next: nil,
	}
	two := &ListNode{
		Val:  2,
		Next: three,
	}
	one := &ListNode{
		Val:  1,
		Next: two,
	}
	head := &ListNode{
		Val:  0,
		Next: one,
	}

	logLinkedList("start", t, head)
	logLinkedList("result", t, reverseLinkedList(head))
}

func logLinkedList(tag string, t *testing.T, head *ListNode) {
	t.Logf("%s: %v", tag, head.Val)
	next := head.Next
	for next != nil {
		t.Logf("%s: %v", tag, next.Val)
		next = next.Next
	}
}
