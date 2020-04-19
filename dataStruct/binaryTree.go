package dataStruct

import "fmt"

type BinaryTreeNode struct {
	Val 	int
	left 	*BinaryTreeNode
	right 	*BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree(v int) *BinaryTree {
	tree := &BinaryTree{}
	tree.Root = &BinaryTreeNode{Val: v}
	return tree
}

func (n *BinaryTreeNode) AddNode(v int) *BinaryTreeNode {
	if n.Val > v {
		if n.left == nil {
			n.left = &BinaryTreeNode{Val: v}
			return n.left
		} else {
			return n.left.AddNode(v)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryTreeNode{Val: v}
			return n.right
		} else {
			return n.right.AddNode(v)
		}
	}
}

type depthNode struct {
	depth int
	node *BinaryTreeNode
}


func (t *BinaryTree) Print() {
	var q []depthNode
	q = append(q, depthNode{depth: 0, node: t.Root})
	currentDepth := 0
	for len(q) > 0 {
		var first = depthNode{}
		first, q = q[0], q[1:]

		if first.depth != currentDepth {
			fmt.Println()
			currentDepth = first.depth
		}
		fmt.Print(first.node.Val, " ")

		if first.node.left != nil {
			q = append(q, depthNode{depth: currentDepth+1, node: first.node.left})
		}
		if first.node.right != nil {
			q = append(q, depthNode{depth: currentDepth+1, node: first.node.right})
		}
	}
}

func (n *BinaryTreeNode) Search(v, cnt int) (bool, int) {
	if n.Val == v {
		return true, cnt
	} else if n.Val > v {
		if n.left != nil {
			return n.left.Search(v, cnt+1)
		}
		return false, cnt
	} else {
		if n.right != nil {
			return n.right.Search(v, cnt+1)
		}
		return false, cnt
	}
}

func (t *BinaryTree) Search(v int) (bool, int) {
	return t.Root.Search(v, 1)
}
