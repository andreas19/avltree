package avltree

import "iter"

// An AVL tree.
type Tree[T any] struct {
	root   *tNode[T]
	cmp    Cmp[T]
	nodups bool
	count  int
}

// New creates a new, empty AVL tree. The function cmp is used for comparing values while
// navigating the tree. If nodups is true, no duplicate values are allowed and Add() will
// return false if the value is already present. Panics if cmp is nil.
func New[T any](cmp Cmp[T], nodups bool) *Tree[T] {
	if cmp == nil {
		panic("cmp cannot be nil")
	}
	return &Tree[T]{cmp: cmp, nodups: nodups}
}

// Collect creates a new AVL tree with values from seq. The function cmp is used for comparing values while
// navigating the tree. If nodups is true, no duplicate values are allowed and Add() will
// return false if the value is already present. Panics if cmp is nil.
func Collect[T any](cmp Cmp[T], nodups bool, seq iter.Seq[T]) *Tree[T] {
	t := New(cmp, nodups)
	for v := range seq {
		t.Add(v)
	}
	return t
}

// Add adds a value to the tree.
func (t *Tree[T]) Add(v T) bool {
	if t.root == nil {
		t.root = &tNode[T]{value: v}
		t.count++
		return true
	} else {
		var b bool
		t.root = t.root.add(v, &b, t.nodups, t.cmp)
		if b {
			t.count++
		}
		return b
	}
}

// Del removes a value from the tree. Returns false if the value was not found.
func (t *Tree[T]) Del(v T) bool {
	if t.root != nil {
		var b bool
		t.root = t.root.remove(v, &b, t.cmp)
		if b {
			t.count--
		}
		return b
	}
	return false
}

// Contains reports whether the value is present.
func (t *Tree[T]) Contains(v T) bool {
	node := t.root
	for node != nil {
		if t.cmp(v, node.value) < 0 {
			node = node.left
		} else if t.cmp(v, node.value) > 0 {
			node = node.right
		} else {
			return true
		}
	}
	return false
}

// Get gets the first value that is equal to v regarding the [Cmp] function.
func (t *Tree[T]) Get(v T) (T, bool) {
	node := t.root
	for node != nil {
		if t.cmp(v, node.value) < 0 {
			node = node.left
		} else if t.cmp(v, node.value) > 0 {
			node = node.right
		} else {
			return node.value, true
		}
	}
	return *new(T), false
}

// GetAll gets all values that are equal to v regarding the [Cmp] function.
func (t *Tree[T]) GetAll(v T) []T {
	result := []T{}
	t.Each(func(x T) {
		if t.cmp(v, x) == 0 {
			result = append(result, x)
		}
	})
	return result
}

// Slice returns a slice with all values from the tree. The values will be sorted
// with regard to the [Cmp] function.
func (t *Tree[T]) Slice() []T {
	sl := []T{}
	if t.root != nil {
		t.root.inorder(func(v T) {
			sl = append(sl, v)
		})
	}
	return sl
}

// Iter returns an iterator over all values from the tree. The values will be sorted
// with regard to the [Cmp] function.
func (t *Tree[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		if t.root != nil {
			if iterTree(t.root, yield) {
				return
			}
		}
	}
}

// returns stop
func iterTree[T any](n *tNode[T], yield func(T) bool) bool {
	if n.left != nil && iterTree(n.left, yield) {
		return true
	}
	if !yield(n.value) {
		return true
	}
	if n.right != nil && iterTree(n.right, yield) {
		return true
	}
	return false
}

// Each traverses the tree in-order and applies fn to each value.
func (t *Tree[T]) Each(fn func(v T)) {
	if t.root == nil {
		return
	}
	t.root.inorder(fn)
}

// IsEmpty returns true if the tree is empty.
func (t *Tree[T]) IsEmpty() bool {
	return t.root == nil
}

// Count returns the count of values in the tree.
func (t *Tree[T]) Count() int {
	return t.count
}

// Clone clones the tree. This method only makes shallow copies of the values.
func (t *Tree[T]) Clone() *Tree[T] {
	var root *tNode[T]
	if t.root != nil {
		root = &tNode[T]{value: t.root.value, height: t.root.height}
		t.root.clone(root)
	}
	return &Tree[T]{root: root, cmp: t.cmp, nodups: t.nodups, count: t.count}
}
