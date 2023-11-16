package main

import (
	"fmt"
	"sync"
)

type T int

type Node struct {
	val   T
	left  *Node
	right *Node
}

func NewTree(v T) *Node {
	tree := &Node{val: v}

	return tree
}

type BinSearchTree struct {
	sync.RWMutex
	root *Node
}

func NewBinSearchTree() *BinSearchTree {
	bst := &BinSearchTree{}

	return bst
}

// Insert inserts the val in the tree
func (bst *BinSearchTree) Insert(val T) {
	bst.Lock()
	defer bst.Unlock()

	node := NewTree(val)
	if bst.root == nil {
		bst.root = node
	} else {
		insert(bst.root, node)
	}
}

// insert internal function to find the correct place for a node in a tree
func insert(root, node *Node) {
	if node.val < root.val {
		if root.left == nil {
			root.left = node
		} else {
			insert(root.left, node)
		}
	} else {
		if root.right == nil {
			root.right = node
		} else {
			insert(root.right, node)
		}
	}
}

func (bst *BinSearchTree) Print() {
	bst.Lock()
	defer bst.Unlock()

	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

func stringify(root *Node, level int) {
	if root != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(root.left, level)
		fmt.Printf(format+"%d\n", root.val)
		stringify(root.right, level)
	}
}

const (
	t1 T = iota + 1
	t2
	t3
	t4
	t5
	t6
	t7
	t8
	t9
	t10
	t11
)

func main() {
	bst := NewBinSearchTree()
	bst.Insert(t3)
	bst.Insert(t2)
	bst.Insert(t4)
	bst.Insert(t9)
	bst.Insert(t5)
	bst.Insert(t6)
	bst.Insert(t10)
	bst.Insert(t1)
	bst.Insert(t7)
	bst.Insert(t8)
	bst.Print()
}

