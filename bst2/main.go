package main

import "fmt"

type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

type BinaryTree struct {
	Root *Node
}

func (t *BinaryTree) insert(value int, node *Node, parent *Node) *Node {
	if node == nil {
		return &Node{Value: value, Parent: parent}
	}
	if value < node.Value {
		node.Left = t.insert(value, node.Left, node)
	} else {
		node.Right = t.insert(value, node.Right, node)
	}
	return node
}

func (t *BinaryTree) Insert(value int) {
	t.Root = t.insert(value, t.Root, nil)
}

func (t *BinaryTree) search(value int, node *Node) *Node {
	if node == nil || node.Value == value {
		return node
	}
	if value < node.Value {
		return t.search(value, node.Left)
	} else {
		return t.search(value, node.Right)
	}
}

func (t *BinaryTree) Search(value int) *Node {
	return t.search(value, t.Root)
}

func main() {
	tree := BinaryTree{}
	values := []int{10, 5, 15, 3, 7, 18}
	for _, v := range values {
		tree.Insert(v)
	}

	result := tree.Search(15)
	if result != nil {
		fmt.Println("Found")
		if result.Parent != nil {
			fmt.Printf("Parent Value: %d\n", result.Parent.Value)
		} else {
			fmt.Println("No Parent")
		}
	} else {
		fmt.Println("Not Found")
	}
}
