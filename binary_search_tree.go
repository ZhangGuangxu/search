package search

import (
	"errors"
	"github.com/ZhangGuangxu/stack"
)

var ErrNotNode = errors.New("not node")

type binarySearchTreeNode struct {
	a int
	left *binarySearchTreeNode
	right *binarySearchTreeNode
}

func newBinarySearchTreeNode(x int) *binarySearchTreeNode {
	return &binarySearchTreeNode{a: x}
}

type BinarySearchTree struct {
	root *binarySearchTreeNode
	nodeCnt int
}

// NewBinarySearchTree creates an instance of BinarySearchTree.
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// IsEmpty returns true if this tree has no node.
func (t *BinarySearchTree) IsEmpty() bool {
	return t.root == nil
}

func (t *BinarySearchTree) Insert(x int) {
	t.nodeCnt++
	n := newBinarySearchTreeNode(x)
	if t.IsEmpty() {
		t.root = n
		return
	}

	node := t.root
	for {
		if x < node.a {
			tmp := node.left
			if tmp == nil {
				node.left = n
				return
			}
			node = tmp
		} else { // x >= node.a
			tmp := node.right
			if tmp == nil {
				node.right = n
				return
			}
			node = tmp
		}
	}
}

func (t *BinarySearchTree) Has(x int) bool {
	node := t.root
	for node != nil {
		if x == node.a {
			return true
		} else if x < node.a {
			node = node.left
		} else { // x >= node.a
			node = node.right
		}
	}
	return false
}

func (t *BinarySearchTree) InorderTraverse() ([]int, error) {
	if t.IsEmpty() {
		return []int{}, nil
	}

	record := make(map[*binarySearchTreeNode]bool, t.nodeCnt)
	r := make([]int, 0, t.nodeCnt)
	sk := stack.NewStackWithSize(t.nodeCnt)
	sk.Push(t.root)
	for !sk.IsEmpty() {
		v, err := sk.Pop()
		if err != nil {
			return []int{}, err
		}
		node, ok := v.(*binarySearchTreeNode)
		if !ok {
			return []int{}, ErrNotNode
		}

		for node != nil {
			_, ok := record[node]
			if ok {
				break
			}
			sk.Push(node)
			node = node.left
		}

		if !sk.IsEmpty() {
			v, err := sk.Pop()
			if err != nil {
				return []int{}, err
			}
			node, ok := v.(*binarySearchTreeNode)
			if !ok {
				return []int{}, ErrNotNode
			}
			r = append(r, node.a)
			record[node] = true
			if node.right != nil {
				sk.Push(node.right)
			}
		}
	}

	return r, nil
}
