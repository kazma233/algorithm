package list

// 反转链表
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
//
// https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/

func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	reversedList := &ListNode{
		Val:  head.Val,
		Next: nil,
	}

	next := head.Next
	for next != nil {
		reversedList = &ListNode{
			Val:  next.Val,
			Next: reversedList,
		}

		next = next.Next
	}

	return reversedList
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode

	curr := head

	for curr != nil {
		next := curr.Next

		// 当前元素的指针指向上一个Node
		curr.Next = prev

		// 两个指针往后移动
		prev = curr
		curr = next
	}

	return prev
}

func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList3(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
