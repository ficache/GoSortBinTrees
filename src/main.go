package main

import (
	"fmt"
)

type Tree struct {
	left  *Tree
	right *Tree
	key   rune
}

func (tree *Tree) insert(aTree *Tree) {
	if aTree.key < tree.key {
		if tree.left != nil {
			tree.left.insert(aTree)
		} else {
			tree.left = aTree
		}
	} else {
		if tree.right != nil {
			tree.right.insert(aTree)
		} else {
			tree.right = aTree
		}
	}
}

func (tree *Tree) traverse(treeVisitor Visitor) {
	if tree.left != nil {
		tree.left.traverse(treeVisitor)
	}

	treeVisitor.visit(tree)

	if tree.right != nil {
		tree.right.traverse(treeVisitor)
	}
}

type Visitor struct{}

func (vis Visitor) visit(node *Tree) {
	fmt.Println(node.key)
}

func main() {
	myTree := Tree{nil, nil, 0}
	slice := []rune{100, 123, 512, 52, 230, 30, 10, 45, 60}
	for _, i := range slice {
		myTree.insert(&Tree{nil, nil, i})
	}

	myTree.traverse(Visitor{})
}
