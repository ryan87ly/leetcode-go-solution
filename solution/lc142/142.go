package lc142

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextN(start *ListNode, n int) (bool, *ListNode) {
	cur := start
	for i := 0; i < n; i++ {
		cur = cur.Next
		if cur == nil {
			return false, nil
		}
	}
	return true, cur
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var slowPtr, fastPtr *ListNode = head, head
	for {
		hasSlow, nextSlow := nextN(slowPtr, 1)
		hasFast, nextFast := nextN(fastPtr, 2)
		if !hasSlow || !hasFast {
			// No loop
			return nil
		}
		slowPtr = nextSlow
		fastPtr = nextFast
		if slowPtr == fastPtr {
			break
		}
	}
	startCycle := head
	for startCycle != slowPtr {
		_, startCycle = nextN(startCycle, 1)
		_, slowPtr = nextN(slowPtr, 1)
	}
	return startCycle
}
