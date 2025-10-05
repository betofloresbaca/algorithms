package structs

import (
	"errors"
	"fmt"
	"strings"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	first *Node[T]
	last  *Node[T]
	count int
}

// func NewLinkedList[T any]() *LinkedList[T] {
// 	return &LinkedList[T]{}
// }

func (ll *LinkedList[T]) Count() int {
	return ll.count
}

func (ll *LinkedList[T]) Push(value T) {
	node := &Node[T]{value: value}
	if ll.count > 0 {
		ll.last.next = node
	} else {
		ll.first = node
	}
	ll.last = node
	ll.count++
}

func (ll *LinkedList[T]) Pop() (T, error) {
	if ll.count == 0 {
		var zero T
		return zero, errors.New("linked list is empty")
	}
	value := ll.last.value
	if ll.count == 1 {
		ll.first = nil
		ll.last = nil
	} else {
		prev := ll.first
		for range ll.count - 2 {
			prev = prev.next
		}
		prev.next = nil
		ll.last = prev
	}
	ll.count--
	return value, nil
}

func (ll *LinkedList[T]) Unshift(value T) {
	node := &Node[T]{value: value}
	node.next = ll.first
	ll.first = node
	if ll.count == 0 {
		ll.last = node
	}
	ll.count++
}

func (ll *LinkedList[T]) Shift() (T, error) {
	if ll.count == 0 {
		var zero T
		return zero, errors.New("linked list is empty")
	}
	value := ll.first.value
	ll.first = ll.first.next
	if ll.first == nil {
		ll.last = nil
	}
	ll.count--
	return value, nil
}

func (ll *LinkedList[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for node := ll.first; node != nil; node = node.next {
		sb.WriteString(fmt.Sprintf("%v", node.value))
		if node.next != nil {
			sb.WriteString(" ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
