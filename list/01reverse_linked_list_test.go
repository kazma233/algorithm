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
