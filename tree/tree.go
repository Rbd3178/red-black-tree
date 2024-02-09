package tree

import "fmt"

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

func (t *tree) GetVal(key int) (string, bool) {
	n := t.root
	for n != nil {
		switch {
		case n.key == key:
			return n.val, true
		case n.key < key:
			n = n.r
		case n.key > key:
			n = n.l
		}
	}
	var v string
	return v, false
}

func (t *tree) ChangeVal(key int, newVal string) bool {
	n := t.root
	for n != nil {
		switch {
		case n.key == key:
			n.val = newVal
			return true
		case n.key < key:
			n = n.r
		case n.key > key:
			n = n.l
		}
	}
	return false
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
	} else {
		for i := 0; i < depth; i++ {
			if i == depth-1 {
				fmt.Print("   |----")
			} else {
				fmt.Print("         ")
			}
		}
		fmt.Print("nil(B)\n")
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

func (t *tree) Insert(key int, val string) bool {
	n := newNode(key, val)
	if t.root == nil {
		n.black = true
		t.root = n
		return true
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
		return false
	}

	if parent.key > key {
		parent.l = n
	} else {
		parent.r = n
	}
	n.p = parent

	t.insertFix(n)

	return true
}
