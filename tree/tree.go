package tree

type keyType int
type valType string

type node struct {
	key keyType
	val valType
	red bool
	l   *node
	r   *node
	p   *node
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

func (t *tree) Search(key keyType) (valType, bool) {
	ptr := t.root
	for ptr != nil {
		switch {
		case ptr.key == key:
			return ptr.val, true
		case ptr.key < key:
			ptr = ptr.r
		case ptr.key > key:
			ptr = ptr.l
		}
	}
	var val valType
	return val, false
}

func (t *tree) Insert(key keyType, val valType) {
	// Decide if duplicates are allowed, if no search first
	n := newNode(key, val)
	n.red = true
}

func (t *tree) Delete(key keyType) {

}

func (t *tree) Prev(key keyType) {

}

func (t *tree) Next(key keyType) {

}
