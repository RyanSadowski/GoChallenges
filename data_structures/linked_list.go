package data_structures

import (
	"strconv"
)

type ListNode struct {
	Value        int
	PreviousNode *ListNode
	NextNode     *ListNode
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func CreateList(node *ListNode) *LinkedList {
	var list = new(LinkedList)
	list.Head = node
	list.Tail = node
	return list
}

func Enqueue(list *LinkedList, node ListNode) {
	list.Tail.NextNode = &node
}

func Dequeue(list *LinkedList) int {
	var tmp = list.Head.Value
	list.Head = list.Head.NextNode
	return tmp
}

func Pop(list *LinkedList) int {
	var tmp = list.Head.Value
	list.Head = list.Head.NextNode
	return tmp
}

func Push(list *LinkedList, node ListNode) *LinkedList {
	list.Head.PreviousNode = &node
	node.NextNode = list.Head.PreviousNode
	list.Head = &node
	return list
}
func IsEmpty(list *LinkedList) bool {
	return list.Head == nil || list.Tail == nil
}

func GetLength(list *LinkedList) int {
	if IsEmpty(list) {
		return 0
	}
	var i = 1
	var currentNode = list.Head

	for currentNode.NextNode != nil {
		i++
		currentNode = currentNode.NextNode
	}
	return i
}

func ToString(list *LinkedList) {
	//var length = getLength(list)
	var strings []string
	var currentNode = list.Head
	var i = 0
	for currentNode.NextNode != nil {
		strings[i] = strconv.Itoa(currentNode.Value)
		currentNode = currentNode.NextNode
		i++
	}
}
