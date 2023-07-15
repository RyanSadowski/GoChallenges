package data_structures_test

import (
	"kek5/data_structures"
	"testing"
)

func TestCreateList(t *testing.T) {
	var node = new(data_structures.ListNode)

	var list = data_structures.CreateList(node)

	if list.Head != node {
		t.Errorf("create node failed to set head")
	}
}

func TestEnqueue(t *testing.T) {
	// Arrange

	// Act

	// Assert
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
