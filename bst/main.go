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

func (n *Node) search(key int) *Node {
	// キーが見つかった場合、現在のノードを返す
	if n == nil || n.Key == key {
		return n
	}
	// キーが現在のノードのキーより小さい場合、左の子ノードを探索
	if key < n.Key {
		return n.Left.search(key)
	}
	// それ以外の場合、右の子ノードを探索
	return n.Right.search(key)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	root := sortedArrayToBST(arr)

	fmt.Println("PreOrder traversal of constructed BST:")
	preOrder(root)
	fmt.Println()

	result := root.search(5)
	if result != nil {
		fmt.Println("Found node with key:", result.Key)
	} else {
		fmt.Println("Key not found")
	}
}
