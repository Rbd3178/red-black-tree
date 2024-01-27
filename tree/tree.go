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

// TODO: use interfaces in node for multiple possible key and val types

func newNode(key keyType, val valType) *node {
	return &node{key: key, val: val}
}

type tree struct {
	root *node
}

func New() *tree {
	return &tree{}
}

func (t *tree) search(key keyType) *node {
	ptr := t.root
	for ptr != nil {
		switch {
		case ptr.key == key:
			return ptr
		case ptr.key < key:
			ptr = ptr.r
		case ptr.key > key:
			ptr = ptr.l
		}
	}
	return ptr
}

func (t *tree) Search(key keyType) (valType, bool) {
	n := t.search(key)
	var v valType
	if n == nil {
		return v, false
	}
	return n.val, true
}

func (t *tree) insert(key keyType, val valType) bool {
	if t.root == nil {
		n := newNode(key, val)
		n.black = true
		t.root = n
		return true
	}
    ptr := t.root
	for (ptr.key )
	n := newNode(key, val)

}

func (t *tree) Delete(key keyType) {

}

func (t *tree) Prev(key keyType) {

}

func (t *tree) Next(key keyType) {

}
