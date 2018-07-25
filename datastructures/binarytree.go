package datastructures

import (
	"fmt"
)

//BinaryTree tree
type BinaryTree struct {
	Root *BinaryNode
}

//NewBinaryTree tree
func NewBinaryTree(value interface{}) *BinaryTree {
	bt := new(BinaryTree)
	bt.Root = &BinaryNode{Value: value}
	return bt
}

//AddLeft add
func (b *BinaryTree) AddLeft(node *BinaryNode, value interface{}) *BinaryNode {
	if node.Left == nil {
		node.Left = &BinaryNode{Value: value}
	}
	return node.Left
}

//AddRight add
func (b *BinaryTree) AddRight(node *BinaryNode, value interface{}) *BinaryNode {
	if node.Right == nil {
		node.Right = &BinaryNode{Value: value}
	}
	return node.Right
}

//Find find
func (b *BinaryTree) Find(value interface{}) *BinaryNode {
	return findNode(b.Root, value)
}

//Infix in
func (b *BinaryTree) Infix() {
	infix(b.Root)
}

func infix(node *BinaryNode) {
	if node == nil {
		return
	}
	infix(node.Left)
	fmt.Printf("%v ", node.Value)
	infix(node.Right)
}

func findNode(node *BinaryNode, value interface{}) *BinaryNode {
	if node == nil {
		return nil
	}
	if node.Value == value {
		return node
	}
	left := findNode(node.Left, value)
	right := findNode(node.Right, value)
	if left != nil {
		return left
	} else {
		return right
	}
}
