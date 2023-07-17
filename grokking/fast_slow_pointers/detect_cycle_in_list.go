package main

func detectCycle(head *EduLinkedListNode) bool {
	// Your code will replce this placeholder return statement
	var slow = head
	var fast = head.next

	for fast != slow {
		if fast == nil || fast.next == nil || fast.next.next == nil {
			return false
		}

		if &fast == &slow {
			return true
		}

		slow = slow.next
		fast = fast.next.next

	}

	return true
}

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}
