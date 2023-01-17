package bst

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

func (r *node) insert(val int) {
	newNode := node{
		data:  val,
		left:  nil,
		right: nil,
	}
	if val < r.data {
		if r.left == nil {
			r.left = &newNode
		} else {
			r.left.insert(val)
		}
	}
	if val > r.data {
		if r.right == nil {
			r.right = &newNode
		} else {
			r.right.insert(val)
		}
	}
}

func (r *node) inorderTraversal() {
	if r == nil {
		return
	}
	r.left.inorderTraversal()
	fmt.Print(r.data, " ")
	r.right.inorderTraversal()
}

func (r *node) preorderTraversal() {
	if r == nil {
		return
	}
	fmt.Print(r.data, " ")
	r.left.preorderTraversal()
	r.right.preorderTraversal()
}

func (r *node) postorderTraversal() {
	if r == nil {
		return
	}
	r.left.postorderTraversal()
	r.right.postorderTraversal()
	fmt.Print(r.data, " ")
}

func Start() {
	fmt.Println("Inside BST Module")
	root := &node{
		data:  10,
		left:  nil,
		right: nil,
	}

	root.insert(50)
	root.insert(40)
	root.insert(20)
	root.insert(30)
	root.insert(60)

	fmt.Println("Inorder Traversal")
	root.inorderTraversal()

	fmt.Println("\nPreorder Traversal")
	root.preorderTraversal()

	fmt.Println("\nPostorder Traversal")
	root.postorderTraversal()

	fmt.Println("")
}
