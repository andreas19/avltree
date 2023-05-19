package avltree

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

// Clone clones the tree.
func (t *Tree[T]) Clone() *Tree[T] {
	var root *tNode[T]
	if t.root != nil {
		root = &tNode[T]{value: t.root.value, height: t.root.height}
		t.root.clone(root)
	}
	return &Tree[T]{root: root, cmp: t.cmp, nodups: t.nodups, count: t.count}
}
