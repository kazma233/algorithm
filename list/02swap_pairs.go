package list

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	one := dummyHead

	for one.Next != nil && one.Next.Next != nil {
		two := one.Next
		three := one.Next.Next

		one.Next = three
		two.Next = three.Next
		three.Next = two

		one = two
	}

	return dummyHead.Next
}

func swapPairs3(one *ListNode) *ListNode {
	if one == nil || one.Next == nil {
		return one
	}

	two := one.Next
	one.Next = swapPairs3(two.Next)
	two.Next = one

	return two
}
