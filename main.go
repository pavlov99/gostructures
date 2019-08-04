package main

import (
	"fmt"
	"strconv"
)

type BinaryTreeNode struct {
	key int
	left *BinaryTreeNode
	right *BinaryTreeNode
}

type BinaryTree struct {
	root *BinaryTreeNode
}

// Inorder treetraversal
func (this *BinaryTreeNode) forEach(callback func(int)) {
	this.left.forEach(callback)
	if this != nil {
		callback(this.key)
	}
	this.right.forEach(callback)
}

// printInorder
func (this *BinaryTreeNode) String() string {
	if this == nil {
		return fmt.Sprintf("{}")
	}
	return "{" + this.left.String() + strconv.Itoa(this.key) + this.right.String() + "}"
}

func (this *BinaryTree) String() string {
	return this.root.String()
}

func (this *BinaryTreeNode) Insert(key int) {
	if key < this.key {
		if this.left == nil {
			this.left = &BinaryTreeNode{key: key}
		} else {
			this.left.Insert(key)
		}
	} else {
		if this.right == nil {
			this.right = &BinaryTreeNode{key: key}
		} else {
			this.right.Insert(key)
		}
	}
}

func (this *BinaryTree) Insert(key int) {
	if this.root == nil {
		this.root = &BinaryTreeNode{key: key}
	} else {
		this.root.Insert(key)
	}
}

func main() {
	fmt.Println("Binary Tree example")
	var t BinaryTree
	fmt.Println(t.String())

	t.Insert(0)
	t.Insert(1)
	t.Insert(2)
	t.Insert(-1)
	fmt.Println(t.String())
}