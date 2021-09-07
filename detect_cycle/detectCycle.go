package detect_cycle

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for slow != p {
				slow = slow.Next
				p = p.Next
			}
			return p
		}
	}
	return nil
}

func TestDetectCycle(t *testing.T) {
	//TODO
}
