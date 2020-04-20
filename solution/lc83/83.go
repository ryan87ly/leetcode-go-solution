package lc83

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	prevNode := head
	for curNode := head.Next; curNode != nil; curNode = curNode.Next {
		if prevNode.Val == curNode.Val {
			prevNode.Next = curNode.Next
		} else {
			prevNode = curNode
		}
	}

	return head
}


