package main

import "fmt"

func main() {
	obj := Constructor()

	obj.AddAtIndex(0, 1)

	fmt.Println(obj.Get(0))

	fmt.Println(obj.Get(1))
}

type Node struct {
	value int
	next  *Node
}

type MyLinkedList struct {
	first *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{nil}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	curNode := this.first

	for i := 0; i < index; i++ {
		if curNode == nil {
			return -1
		}
		curNode = curNode.next
	}

	if curNode == nil {
		return -1
	}

	return curNode.value
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := Node{val, nil}

	if this.first == nil {
		this.first = &newNode
	} else {
		newNode.next = this.first
		this.first = &newNode
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	var newNode Node
	newNode.value = val

	curNode := this.first

	if curNode == nil {
		this.first = &newNode
		return
	}

	for curNode.next != nil {
		curNode = curNode.next
	}

	curNode.next = &newNode
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	newNode := Node{val, nil}

	if index == 0 {
		newNode.next = this.first
		this.first = &newNode
		return
	}

	curNode := this.first

	for i := 0; i < index-1; i++ {
		if curNode == nil {
			return
		}
		curNode = curNode.next
	}

	if curNode == nil {
		return
	}

	newNode.next = curNode.next
	curNode.next = &newNode
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.first = this.first.next
		return
	}

	curNode := this.first

	for i := 0; i < index-1; i++ {
		if curNode == nil {
			return
		}
		curNode = curNode.next
	}

	if curNode.next == nil {
		return
	}

	curNode.next = curNode.next.next
}
