package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func sortedArrayToBST(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	mid := len(arr) / 2

	node := &Node{Key: arr[mid]}
	node.Left = sortedArrayToBST(arr[:mid])
	node.Right = sortedArrayToBST(arr[mid+1:])

	return node
}

func preOrder(n *Node) {
	if n == nil {
		return
	}

	preOrder(n.Left)
	fmt.Print(n.Key, " ")
	preOrder(n.Right)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	root := sortedArrayToBST(arr)

	fmt.Println("PreOrder traversal of constructed BST:")
	preOrder(root)
}
