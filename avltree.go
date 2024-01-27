// Package avltree implements a generic AVL tree.
package avltree

import "cmp"

// A Cmp function compares two values and returns 0 if a == b, -1 if a < b, and 1 if a > b.
type Cmp[T any] func(a, b T) int

// CmpOrd is an implementation of the [Cmp] type for ordered values.
func CmpOrd[T cmp.Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
