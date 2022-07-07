package main

import "fmt"

type Node struct{
	data string
	left *Node
	right *Node
}
func Inorder(root *Node) {
	if root == nil {
		return
	}

	Inorder(root.left)
	fmt.Printf(root.data)
	Inorder(root.right)
}

func Preorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf(root.data)
	Preorder(root.left)
	Preorder(root.right)
}

func Postorder(root *Node) {
	if root == nil {
		return
	}

	Postorder(root.left)
	Postorder(root.right)
	fmt.Printf(root.data)
}

func main(){
	root := &Node{"+", nil, nil}
	root.left = &Node{"a", nil, nil}
	root.right = &Node{"-", nil, nil}
	root.right.left = &Node{"b", nil, nil}
	root.right.right = &Node{"c", nil, nil}

	fmt.Println("Pre Order Traversal: ")
	Preorder(root)
	fmt.Println("\n")
	fmt.Println("Post Order Traversal: ")
	Postorder(root)
	fmt.Println("\n")
	fmt.Println("InOrder Traversal: ")
	Inorder(root)
}