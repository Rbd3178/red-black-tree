package tree

import (
	"errors"
	// "fmt"
	"golang.org/x/exp/constraints"
)

type node[kT constraints.Ordered, vT any] struct {
	key   kT
	val   vT
	black bool
	l     *node[kT, vT]
	r     *node[kT, vT]
	p     *node[kT, vT]
}

func newNode[kT constraints.Ordered, vT any](key kT, val vT) *node[kT, vT] {
	return &node[kT, vT]{key: key, val: val}
}

type Tree[kT constraints.Ordered, vT any] struct {
	root *node[kT, vT]
	size int
	max *node[kT, vT]
	min *node[kT, vT]
}

func (t *Tree[kT, vT]) getNode(key kT) *node[kT, vT] {
	n := t.root
	for n != nil && n.key != key {
		if n.key < key {
			n = n.r
		} else {
			n = n.l
		}
	}
	return n
}

func (t *Tree[kT, vT]) At(key kT) (vT, error) {
	n := t.getNode(key)
	if n == nil {
		var v vT
		return v, errors.New("key doesn't exist")
	}
	return n.val, nil
}

func (t *Tree[kT, vT]) Assign(key kT, val vT) error {
	n := t.getNode(key)
	if n == nil {
		return errors.New("key doesn't exist")
	}
	n.val = val
	return nil
}

func (t *Tree[kT, vT]) InOrder() [][]any {
	var out [][]any
	var s []*node[kT, vT]
	n := t.root
	for n != nil || len(s) != 0 {
		if n == nil && len(s) != 0 {
			var pair = []any{s[len(s)-1].key, s[len(s)-1].val}
			out = append(out, pair)
			n = s[len(s)-1].r
			s = s[:len(s)-1]
		} else {
			s = append(s, n)
			n = n.l
		}
	}
	return out
}

func (t *Tree[kT, vT]) Size() int {
	return t.size
}

func (t *Tree[kT, vT]) Max() (kT, vT, error) {
	if t.root == nil {
		var k kT
		var v vT
		return k, v, errors.New("tree is empty")
	}
	return t.max.key, t.max.val, nil
}

func (t *Tree[kT, vT]) Min() (kT, vT, error) {
	if t.root == nil {
		var k kT
		var v vT
		return k, v, errors.New("tree is empty")
	}
	return t.min.key, t.min.val, nil
}

func (t *Tree[kT, vT]) Next(key kT) (kT, vT, error) {
	n := t.root
	var curBest *node[kT, vT]
	if n == nil {
		var k kT
		var v vT
		return k, v, errors.New("larger key doesn't exist (tree is empty)")
	}
	for (n.key <= key && n.r != nil) || (n.key > key && n.l != nil) {
		if n.key <= key && n.r != nil {
			n = n.r
		} else {
			curBest = n
			n = n.l
		}
	}
	if n.key <= key && curBest == nil {
		var k kT
		var v vT
		return k, v, errors.New("larger key doesn't exist")
	}
	if n.key > key {
		return n.key, n.val, nil
	}
	return curBest.key, curBest.val, nil
}

func (t *Tree[kT, vT]) Prev(key kT) (kT, vT, error) {
	n := t.root
	var curBest *node[kT, vT]
	if n == nil {
		var k kT
		var v vT
		return k, v, errors.New("smaller key doesn't exist (tree is empty)")
	}
	for (n.key >= key && n.l != nil) || (n.key < key && n.r != nil) {
		if n.key >= key && n.l != nil {
			n = n.l
		} else {
			curBest = n
			n = n.r
		}
	}
	if n.key >= key && curBest == nil {
		var k kT
		var v vT
		return k, v, errors.New("smaller key doesn't exist")
	}
	if n.key < key {
		return n.key, n.val, nil
	}
	return curBest.key, curBest.val, nil
}

func (t *Tree[keyType, valType]) leftRotate(n *node[keyType, valType]) {
	child := n.r

	n.r = child.l
	if child.l != nil {
		child.l.p = n
	}

	child.p = n.p
	switch {
	case n == t.root:
		t.root = child
	case n == n.p.l:
		n.p.l = child
	case n == n.p.r:
		n.p.r = child
	}

	child.l = n
	n.p = child
}

func (t *Tree[keyType, valType]) rightRotate(n *node[keyType, valType]) {
	child := n.l

	n.l = child.r
	if child.r != nil {
		child.r.p = n
	}

	child.p = n.p
	switch {
	case n == t.root:
		t.root = child
	case n == n.p.l:
		n.p.l = child
	case n == n.p.r:
		n.p.r = child
	}

	child.r = n
	n.p = child
}

func (t *Tree[keyType, valType]) insertFix(n *node[keyType, valType]) {
	for n != t.root && !n.p.black {
		parent := n.p
		grand := n.p.p
		switch {
		case parent == grand.l && n == parent.r && (grand.r == nil || grand.r.black):
			t.leftRotate(parent)
			t.rightRotate(grand)
			grand.black = false
			n.black = true
			return
		case parent == grand.r && n == parent.l && (grand.l == nil || grand.l.black):
			t.rightRotate(parent)
			t.leftRotate(grand)
			grand.black = false
			n.black = true
			return
		case parent == grand.l && n == parent.l && (grand.r == nil || grand.r.black):
			t.rightRotate(grand)
			grand.black = false
			parent.black = true
			return
		case parent == grand.r && n == parent.r && (grand.l == nil || grand.l.black):
			t.leftRotate(grand)
			grand.black = false
			parent.black = true
			return
		case parent == grand.l && !grand.r.black:
			parent.black = true
			grand.r.black = true
			grand.black = false
			n = grand
		case parent == grand.r && !grand.l.black:
			parent.black = true
			grand.l.black = true
			grand.black = false
			n = grand
		}
	}

	if n == t.root {
		n.black = true
	}
}

func (t *Tree[keyType, valType]) Insert(key keyType, val valType) error {
	n := newNode(key, val)

	if t.root == nil {
		n.black = true
		t.root = n
		t.size++
		t.max = n
		t.min = n
		return nil
	}

	parent := t.root
	for parent.key > key && parent.l != nil || parent.key < key && parent.r != nil {
		if parent.key < key {
			parent = parent.r
		} else {
			parent = parent.l
		}
	}
	if parent.key == key {
		return errors.New("key already exists")
	}

	if parent.key > key {
		parent.l = n
	} else {
		parent.r = n
	}
	n.p = parent

	t.insertFix(n)

	t.size++
	if key > t.max.key {
		t.max = n
	}
	if key < t.min.key {
		t.min = n
	}
	return nil
}

/*
func (t *Tree) Delete(key int) error {
	n := t.root

	for n != nil && (n.key > key && n.l != nil || n.key < key && n.r != nil) {
		if n.key < key {
			n = n.r
		} else {
			n = n.l
		}
	}
	if n == nil {
		return errors.New("key doesn't exist")
	}

	if n == n.p.l {
		switch {
		case n.l != nil && n.r != nil:
			// something
		case n.l != nil:
			n.l.p = n.p
			n.p.l = n.l
		case n.r != nil:
			n.r.p = n.p
			n.p.l = n.r
		default:
			n.p.l = nil
		}
	} else {
		switch {
		case n.l != nil && n.r != nil:
			// something
		case n.l != nil:
			n.l.p = n.p
			n.p.r = n.l
		case n.r != nil:
			n.r.p = n.p
			n.p.r = n.r
		default:
			n.p.r = nil
		}
	}

	return nil
}
*/
