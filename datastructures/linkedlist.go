package datastructures

import (
	"fmt"
)

//List list
type List struct {
	head *Node
	size int
}

//NewList new
func NewList() *List {
	return new(List)
}

//IsEmpty empty
func (l *List) IsEmpty() bool {
	return l.size == 0
}

//Add add
func (l *List) Add(val interface{}) {
	newNode := &Node{Value: val, Next: nil}
	lastNode := l.Last()
	if lastNode == nil {
		l.head = newNode
	} else {
		lastNode.Next = newNode
		newNode.Prev = lastNode
	}
	l.size++
}

//Last last
func (l *List) Last() *Node {
	if l.IsEmpty() {
		return nil
	}
	var n *Node
	for n = l.head; n.Next != nil; n = n.Next { //why = works over :=
	}
	return n
}

//Find find
func (l *List) Find(val interface{}) *Node {
	if l.IsEmpty() {
		return nil
	}
	var n *Node
	for n = l.head; n != nil; n = n.Next {
		if n.Value == val {
			return n
		}
	}
	return nil
}

//Display display
func (l *List) Display() {
	for n := l.head; n != nil; n = n.Next {
		fmt.Printf("%d ", n.Value)
	}
	fmt.Println()
}
