package avltree

type tNode[T any] struct {
	value       T
	left, right *tNode[T]
	height      int
}

func (n *tNode[T]) computeHeight() {
	height := -1
	if n.left != nil {
		height = maxInt(height, n.left.height)
	}
	if n.right != nil {
		height = maxInt(height, n.right.height)
	}
	n.height = height + 1
}

func (n *tNode[T]) heightDiff() int {
	var heightLeft, heightRight int
	if n.left != nil {
		heightLeft = 1 + n.left.height
	}
	if n.right != nil {
		heightRight = 1 + n.right.height
	}
	return heightLeft - heightRight
}

func (n *tNode[T]) add(v T, b *bool, nodups bool, cmp Cmp[T]) *tNode[T] {
	if nodups && cmp(v, n.value) == 0 {
		return n
	}
	newRoot := n
	if cmp(v, n.value) <= 0 {
		n.left = n.addToSubTree(n.left, v, b, nodups, cmp)
		if n.heightDiff() == 2 {
			if cmp(v, n.left.value) <= 0 {
				newRoot = n.rotateRight()
			} else {
				newRoot = n.rotateLeftRight()
			}
		}
	} else {
		n.right = n.addToSubTree(n.right, v, b, nodups, cmp)
		if n.heightDiff() == -2 {
			if cmp(v, n.right.value) > 0 {
				newRoot = n.rotateLeft()
			} else {
				newRoot = n.rotateRightLeft()
			}
		}
	}
	newRoot.computeHeight()
	return newRoot
}

func (n *tNode[T]) addToSubTree(parent *tNode[T], v T, b *bool, nodups bool, cmp Cmp[T]) *tNode[T] {
	if parent == nil {
		*b = true
		return &tNode[T]{value: v}
	}
	return parent.add(v, b, nodups, cmp)
}

func (n *tNode[T]) rotateRight() *tNode[T] {
	newRoot := n.left
	grand := newRoot.right
	n.left = grand
	newRoot.right = n
	n.computeHeight()
	return newRoot
}

func (n *tNode[T]) rotateRightLeft() *tNode[T] {
	child := n.right
	newRoot := child.left
	grand1 := newRoot.left
	grand2 := newRoot.right
	child.left = grand2
	n.right = grand1
	newRoot.left = n
	newRoot.right = child
	child.computeHeight()
	n.computeHeight()
	return newRoot
}

func (n *tNode[T]) rotateLeft() *tNode[T] {
	newRoot := n.right
	grand := newRoot.left
	n.right = grand
	newRoot.left = n
	n.computeHeight()
	return newRoot
}

func (n *tNode[T]) rotateLeftRight() *tNode[T] {
	child := n.left
	newRoot := child.right
	grand1 := newRoot.left
	grand2 := newRoot.right
	child.right = grand1
	n.left = grand2
	newRoot.left = child
	newRoot.right = n
	child.computeHeight()
	n.computeHeight()
	return newRoot
}

func (n *tNode[T]) removeFromParent(parent *tNode[T], v T, b *bool, cmp Cmp[T]) *tNode[T] {
	if parent != nil {
		return parent.remove(v, b, cmp)
	}
	return nil
}

func (n *tNode[T]) remove(v T, b *bool, cmp Cmp[T]) *tNode[T] {
	newRoot := n
	if cmp(v, n.value) == 0 {
		*b = true
		if n.left == nil {
			return n.right
		}
		child := n.left
		for child.right != nil {
			child = child.right
		}
		childValue := child.value
		n.left = n.removeFromParent(n.left, childValue, b, cmp)
		n.value = childValue
		if n.heightDiff() == -2 {
			if n.right.heightDiff() <= 0 {
				newRoot = n.rotateLeft()
			} else {
				newRoot = n.rotateRightLeft()
			}
		}
	} else if cmp(v, n.value) < 0 {
		n.left = n.removeFromParent(n.left, v, b, cmp)
		if n.heightDiff() == -2 {
			if n.right.heightDiff() <= 0 {
				newRoot = n.rotateLeft()
			} else {
				newRoot = n.rotateRightLeft()
			}
		}
	} else {
		n.right = n.removeFromParent(n.right, v, b, cmp)
		if n.heightDiff() == 2 {
			if n.left.heightDiff() >= 0 {
				newRoot = n.rotateRight()
			} else {
				newRoot = n.rotateLeftRight()
			}
		}
	}
	newRoot.computeHeight()
	return newRoot
}

func (n *tNode[T]) inorder(fn func(v T)) {
	if n.left != nil {
		n.left.inorder(fn)
	}
	fn(n.value)
	if n.right != nil {
		n.right.inorder(fn)
	}
}

func (n *tNode[T]) clone(c *tNode[T]) {
	if n.left != nil {
		c.left = &tNode[T]{value: n.left.value, height: n.left.height}
		n.left.clone(c.left)
	}
	if n.right != nil {
		c.right = &tNode[T]{value: n.right.value, height: n.right.height}
		n.right.clone(c.right)
	}
}
