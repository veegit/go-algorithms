package datastructures

//Node node
type Node struct {
	Next  *Node
	Prev  *Node
	Value interface{}
}

//BinaryNode node
type BinaryNode struct {
	Left  *BinaryNode
	Right *BinaryNode
	Value interface{}
}

//TreeNode node
type TreeNode struct {
	Children []*TreeNode
	Value    interface{}
}
