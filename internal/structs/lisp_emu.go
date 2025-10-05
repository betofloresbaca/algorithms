package structs

import (
	"fmt"
	"strings"
)

type cell[T any] struct {
	car T
	cdr *cell[T]
}

func Cons[T any](car T, cdr *cell[T]) *cell[T] {
	return &cell[T]{car: car, cdr: cdr}
}

func Car[T any](c *cell[T]) T {
	if c == nil {
		var zero T
		return zero
	}
	return c.car
}

func Cdr[T any](c *cell[T]) *cell[T] {
	if c == nil {
		return nil
	}
	return c.cdr
}

func (c *cell[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for node := c; node != nil; node = node.cdr {
		sb.WriteString(fmt.Sprintf("%v", node.car))
		if node.cdr != nil {
			sb.WriteString(" ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
