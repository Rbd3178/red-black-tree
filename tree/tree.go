package tree

type keyType int
type valType string

type node struct {
	key   keyType
	val   valType
	black bool
	l     *node
	r     *node
	p     *node
}

func newNode(key keyType, val valType) *node {
	return &node{key: key, val: val}
}

type tree struct {
	root *node
}

func New() *tree {
	return &tree{}
}

func (t *tree) Search(key keyType) (valType, bool) {
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
	var v valType
	return v, false
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
			n.black = true
			return
		case parent == grand.r && n == parent.r && (grand.l == nil || grand.l.black):
			t.leftRotate(grand)
			grand.black = false
			n.black = true
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

func (t *tree) Insert(key keyType, val valType) bool {
	n := newNode(key, val)
	if t.root == nil {
		n.black = true
		t.root = n
		return true
	}

	parent := t.root
	for parent.key > key && parent.l != nil || parent.key < key && parent.r != nil {
		switch {
		case parent.key == key:
			return false
		case parent.key < key:
			parent = parent.r
		case parent.key > key:
			parent = parent.l
		}
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