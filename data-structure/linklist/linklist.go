package linklist

import (
	"fmt"
)

//Node Link list Node
type Node struct {
	Data int
	Next *Node
}

//CreateEmptyLinkList will create the empty link list
func CreateEmptyLinkList() *Node {
	return nil
}

//Append will be used to append data in the link list
func (head *Node) Append(data int) *Node {
	newNode := &Node{Data: data}
	last := head
	if head == nil {
		head = newNode
		return newNode
	}
	for last.Next != nil {
		last = last.Next
	}
	last.Next = newNode
	return head
}

//AddInFront will be used to add in front
func (head *Node) AddInFront(data int) *Node {
	newNode := &Node{Data: data}
	newNode.Next = head
	return newNode
}

//PrintList will be used to print the link list
func (head *Node) PrintList() {
	node := head
	for node != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
	return
}

//ReverseLinkList will be used tot reverse the link list
func (head *Node) ReverseLinkList() *Node {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	var prev, next *Node
	current := head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	head = prev
	return head
}

//RecursiveReverse will be used to reverse link list recursively
func (head *Node) RecursiveReverse() *Node {
	var first, rest *Node

	if head == nil {
		return nil
	}
	first = head
	rest = first.Next
	if rest == nil {
		return first
	}
	rest.RecursiveReverse()
	first.Next.Next = first
	first.Next = nil
	head = rest
	return head
}
