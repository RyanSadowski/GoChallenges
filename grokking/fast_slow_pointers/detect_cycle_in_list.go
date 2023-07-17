package main

func detectCycle(head *Node) bool {
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

type Node struct {
	data int
	next *Node
}
