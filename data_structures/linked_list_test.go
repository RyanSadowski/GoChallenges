package data_structures_test

import (
	"kek5/data_structures"
	"testing"
)

func CreateNode(t *testing.T) {
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
