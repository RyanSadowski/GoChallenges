package data_structures_test

import (
	"kek5/data_structures"
	"testing"
)

func TestCreateNode(t *testing.T) {
	var node = data_structures.CreateNode(5)

	if node.NextNode != nil {
		t.Errorf("new node should have null NextNode property")
	}

	if node.PreviousNode != nil {
		t.Errorf("new node should have null PreviousNode property")
	}

	if node.Value != 5 {
		t.Errorf("why tho")
	}

}

func TestCreateList(t *testing.T) {
	var node = new(data_structures.ListNode)

	var list = data_structures.CreateList(node)
	if list.Tail != node {
		t.Errorf("create node failed to set tail")
	}

	if list.Head != node {
		t.Errorf("create node failed to set head")
	}
}

func TestEnqueue(t *testing.T) {
	// Arrange
	var node = new(data_structures.ListNode)
	node.Value = 1
	var list = data_structures.CreateList(node)
	var aNode = data_structures.CreateNode(2)
	var newNode = data_structures.CreateNode(7)

	// Act
	data_structures.Enqueue(list, aNode)
	data_structures.Enqueue(list, newNode)
	var length = data_structures.GetLength(list)
	var expected = 3

	// Assert

	if length != expected {
		t.Errorf("Length after Enqueue is wrong, Found %d, expected %d.", length, expected)
	}
}

func TestDequeue(t *testing.T) {
	// Arrange
	var node = data_structures.CreateNode(1)
	var node2 = data_structures.CreateNode(2)
	var node3 = data_structures.CreateNode(3)
	var node4 = data_structures.CreateNode(4)
	var node5 = data_structures.CreateNode(5)
	var list = data_structures.CreateList(node)

	// Act
	data_structures.Enqueue(list, node2)
	if data_structures.GetLength(list) != 2 {
		// error
	}

	data_structures.Enqueue(list, node3)
	data_structures.Enqueue(list, node4)
	data_structures.Enqueue(list, node5)

	if data_structures.GetLength(list) != 5 {
		// error
	}

	var firstD = data_structures.Dequeue(list)
	var secondD = data_structures.Dequeue(list)
	data_structures.Dequeue(list)
	data_structures.Dequeue(list)
	var fifthD = data_structures.Dequeue(list)

	if firstD != node.Value {
		t.Errorf("first DeQueue is bad")
	}
	if secondD != node2.Value {
		t.Errorf("2nd DeQueue is bad")
	}

	if fifthD != node5.Value {
		t.Errorf("fifth DeQueue is bad")
	}

	if data_structures.GetLength(list) != 0 {
		t.Errorf("list didn't shrink during dequeue")
	}

}

func TestLength(t *testing.T) {

	// Arrange
	var list = new(data_structures.LinkedList)
	var headNode = new(data_structures.ListNode)
	var tailNode = new(data_structures.ListNode)

	headNode.NextNode = tailNode
	tailNode.PreviousNode = headNode
	list.Head = headNode
	list.Tail = tailNode

	// Act
	var length = data_structures.GetLength(list)
	var expected = 2

	// Assert
	if length != expected {
		t.Errorf("GetLength returned %d but we're expecting %d", length, expected)
	}

}
