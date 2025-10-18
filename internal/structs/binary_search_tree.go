package structs

import "cmp"

type BstNode[T cmp.Ordered] struct {
	value T
	left  *BstNode[T]
	right *BstNode[T]
}

func NewBstNode[T cmp.Ordered](value T) *BstNode[T] {
	return &BstNode[T]{
		value: value,
	}
}

func (node *BstNode[T]) Add(value T) *BstNode[T] {
	if value < node.value {
		if node.left == nil {
			node.left = NewBstNode(value)
		} else {
			node.left.Add(value)
		}
	} else {
		if node.right == nil {
			node.right = NewBstNode(value)
		} else {
			node.right.Add(value)
		}
	}
	return node
}
