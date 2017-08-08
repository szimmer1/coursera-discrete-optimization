package main

import (
	"errors"
	"fmt"
)

type BTNode struct {
	v     int
	left  *BTNode
	right *BTNode
	root  *BTNode
}

func SerializeTree(root *BTNode) string {
	if root == nil {
		return "nil"
	}
	return fmt.Sprintf(
		"%d(%s, %s)",
		root.v,
		SerializeTree(root.left),
		SerializeTree(root.right),
	)
}

func Inorder(root *BTNode, f func(*BTNode)) {
	if root != nil {
		Inorder(root.left, f)
		f(root)
		Inorder(root.right, f)
	}
}

func SerializeTreeInorder(root *BTNode) string {
	if root == nil {
		return ""
	}
	s := SerializeTreeInorder(root.left)
	s += fmt.Sprintf(" %d", root.v)
	return s + SerializeTreeInorder(root.right)
}

type SparseM interface {
	Get(r int, c int) (int, error)
	Set(r int, c int, v int) error
	Serialize() string
}

type sparseM struct {
	root *BTNode
}

func (m *sparseM) Get(r int, c int) (int, error) {
	rowNode, _ := getNode(m.root, r)
	if rowNode == nil {
		return 0, errors.New("row not found")
	}
	colNode, _ := getNode(rowNode, c)
	if colNode == nil {
		return 0, errors.New("col not found")
	}
	return colNode.v, nil
}

func (m *sparseM) Set(r int, c int, v int) error {
	// if this is the first matrix value, set it as row root
	if m.root == nil {
		m.root = &BTNode{
			v: r,
			root: &BTNode{
				v: c,
			},
		}
		return nil
	}

	rowNode, rowNodeP := getNode(m.root, r)

	// if first time we have encountered this row index,
	// insert the row node correctly, with column root
	if rowNode == nil {
		return insertNextNode(rowNodeP, &BTNode{
			v: r,
			root: &BTNode{
				v: c,
			},
		})
	}

	colNode, colNodeP := getNode(rowNode, c)

	// if first time we have encountered this col index,
	// insert the col node correctly
	if colNode == nil {
		return insertNextNode(colNodeP, &BTNode{
			v: c,
		})
	}

	// if this col node exists, override its value
	colNode.v = v

	return nil
}

func getNode(root *BTNode, v int) (*BTNode, *BTNode) {
	if root == nil {
		return nil, nil
	}
	var c, p *BTNode
	c = root
	for c != nil {
		if c.v == v {
			return c, p
		}
		p = c
		if c.v < v {
			c = c.right
		} else {
			c = c.left
		}
	}
	return nil, p
}

func insertNextNode(root *BTNode, n *BTNode) error {
	if root == nil || n == nil {
		return errors.New("root or n is nil")
	}
	if root.v < n.v {
		root.right = n
	} else {
		root.left = n
	}
	return nil
}
