package tree

import (
	"errors"
	"fmt"
)

type node struct {
	key   int
	val   string
	black bool
	l     *node
	r     *node
	p     *node
}

func newNode(key int, val string) *node {
	return &node{key: key, val: val}
}

type tree struct {
	root *node
}

func New() *tree {
	return &tree{}
}

func (t *tree) getNode(key int) *node {
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

func (t *tree) At(key int) (string, error) {
	n := t.getNode(key)
	if n == nil {
		var v string
		return v, errors.New("key doesn't exist")
	}
	return n.val, nil
}

func (t *tree) Assign(key int, val string) error {
	n := t.getNode(key)
	if n == nil {
		return errors.New("key doesn't exist")
	}
	n.val = val
	return nil
}

/*func nextNode(n *node) *node {
	if n.r != nil {
		n = n.r
		for n.l != nil {
			n = n.l
		}
		return n
	}
	for n.p != nil && n == n.p.r {
		n = n.p
	}
	return n.p
}

func prevNode(n *node) *node {
	if n.l != nil {
		n = n.l
		for n.r != nil {
			n = n.r
		}
		return n
	}
	for n.p != nil && n == n.p.l {
		n = n.p
	}
	return n.p
}*/

func (t *tree) Next(key int) (int, string, error) {
	n := t.root
	if n == nil {
		var k int
		var v string
		return k, v, errors.New("larger key doesn't exist (tree is empty)")
	}
	for (n.key <= key && n.r != nil) || (n.key > key && n.l != nil) {
		if n.key <= key && n.r != nil {
			n = n.r
		} else {
			n = n.l
		}
	}
	if n.key <= key {
		var k int
		var v string
		return k, v, errors.New("larger key doesn't exist")
	}
	return n.key, n.val, nil
}

func (t *tree) Prev(key int) (int, string, error) {
	n := t.root
	if n == nil {
		var k int
		var v string
		return k, v, errors.New("smaller key doesn't exist (tree is empty)")
	}
	for (n.key >= key && n.l != nil) || (n.key < key && n.r != nil) {
		if n.key >= key && n.l != nil {
			n = n.l
		} else {
			n = n.r
		}
	}
	if n.key >= key {
		var k int
		var v string
		return k, v, errors.New("smaller key doesn't exist")
	}
	return n.key, n.val, nil
}

func visualizeInternal(n *node, depth int) {
	if n != nil {
		visualizeInternal(n.r, depth+1)
		for i := 0; i < depth; i++ {
			if i == depth-1 {
				fmt.Print("   |----")
			} else {
				fmt.Print("        ")
			}
		}
		fmt.Print(n.key)
		if n.black {
			fmt.Print("(B)\n")
		} else {
			fmt.Print("(R)\n")
		}
		visualizeInternal(n.l, depth+1)
	}
}

func (t *tree) Visualize() {
	visualizeInternal(t.root, 0)
}

func (t *tree) leftRotate(n *node) {
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

func (t *tree) rightRotate(n *node) {
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

func (t *tree) insertFix(n *node) {
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

func (t *tree) Insert(key int, val string) error {
	n := newNode(key, val)

	if t.root == nil {
		n.black = true
		t.root = n
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

	return nil
}

func (t *tree) Delete(key int) error {
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
