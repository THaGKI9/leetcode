package problem0203

import "github.com/thagki9/leetcode/kit"

/*
Source: https://leetcode.com/problems/remove-linked-list-elements
Test Case:
[1,2,6,3,4,5,6]
6
*/

// ListNode is a list node
type ListNode = kit.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	virtualHead := &ListNode{Next: head}
	head = virtualHead

	for head.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	return virtualHead.Next
}
