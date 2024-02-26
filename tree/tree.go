package tree

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
	"strings"
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

// Tree represents an associative container implemented as a red-black tree.
// Red-black trees are self-balancing binary search trees that provide efficient
// operations for insertion, deletion, and search,
// which all have a complexity of O(log(n)), where n is the amount of key-value pairs stored.
// Additionally, the tree has the ability to return its elements in order.
// Use the provided methods to interact with the tree.
type Tree[kT constraints.Ordered, vT any] struct {
	root *node[kT, vT]
	size int
	max  *node[kT, vT]
	min  *node[kT, vT]
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

// Returns a value stored by the provided key.
// Returns an error if the provided key is not in the structure.
// O(log(n))
func (t *Tree[kT, vT]) At(key kT) (vT, error) {
	n := t.getNode(key)
	if n == nil {
		var v vT
		return v, errors.New("key doesn't exist")
	}
	return n.val, nil
}

// Changes the value stored by the provided key to the one provided.
// Returns an error if the provided key is not in the structure.
// O(log(n))
func (t *Tree[kT, vT]) Assign(key kT, val vT) error {
	n := t.getNode(key)
	if n == nil {
		return errors.New("key doesn't exist")
	}
	n.val = val
	return nil
}

// Returns a slice of all key and value pairs,
// sorted by key in ascending order.
// O(n)
func (t *Tree[kT, vT]) InOrder() [][]any {
	var out [][]any
	var s []*node[kT, vT]
	n := t.root
	for n != nil || len(s) != 0 {
		if n == nil && len(s) != 0 {
			out = append(out, []any{s[len(s)-1].key, s[len(s)-1].val})
			n = s[len(s)-1].r
			s = s[:len(s)-1]
		} else {
			s = append(s, n)
			n = n.l
		}
	}
	return out
}

// Returns the amount of nodes in the tree.
// O(1)
func (t *Tree[kT, vT]) Size() int {
	return t.size
}

// Returns the key and value pair with the biggest key.
// Returns an error if the tree is empty.
// O(1)
func (t *Tree[kT, vT]) Max() (kT, vT, error) {
	if t.root == nil {
		var k kT
		var v vT
		return k, v, errors.New("tree is empty")
	}
	return t.max.key, t.max.val, nil
}

// Returns the key and value pair with the smallest key.
// Returns an error if the tree is empty.
// O(1)
func (t *Tree[kT, vT]) Min() (kT, vT, error) {
	if t.root == nil {
		var k kT
		var v vT
		return k, v, errors.New("tree is empty")
	}
	return t.min.key, t.min.val, nil
}

// Returns the first key and value pair,
// where the key is larger than the one provided.
// Returns an error if there is no larger key in the tree.
// O(log(n))
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

// Returns the first key and value pair,
// where the key is smaller than the one provided.
// Returns an error if there is no smaller key in the tree.
// O(log(n))
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

func (t *Tree[kT, vT]) leftRotate(n *node[kT, vT]) {
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

func (t *Tree[kT, vT]) rightRotate(n *node[kT, vT]) {
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

func (n *node[kT, vT]) isBlack() bool {
	return n == nil || n.black
}

// Adds a node with provided key and value to the tree.
// Returns an error if the key is already in the structure.
// O(log(n))
func (t *Tree[kT, vT]) Insert(key kT, val vT) error {
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

func (t *Tree[keyType, valType]) insertFix(n *node[keyType, valType]) {
	for n != t.root && !n.p.black {
		par := n.p
		gran := n.p.p
		switch {
		case par == gran.l && n == par.r && gran.r.isBlack():
			t.leftRotate(par)
			t.rightRotate(gran)
			gran.black = false
			n.black = true
			return
		case par == gran.r && n == par.l && gran.l.isBlack():
			t.rightRotate(par)
			t.leftRotate(gran)
			gran.black = false
			n.black = true
			return
		case par == gran.l && n == par.l && gran.r.isBlack():
			t.rightRotate(gran)
			gran.black = false
			par.black = true
			return
		case par == gran.r && n == par.r && gran.l.isBlack():
			t.leftRotate(gran)
			gran.black = false
			par.black = true
			return
		case par == gran.l && !gran.r.isBlack():
			par.black = true
			gran.r.black = true
			gran.black = false
			n = gran
		case par == gran.r && !gran.l.isBlack():
			par.black = true
			gran.l.black = true
			gran.black = false
			n = gran
		}
	}

	if n == t.root {
		n.black = true
	}
}

// Deletes the node with the provided key from the tree.
// Returns an error if the key is not found in the structure.
// O(log(n))
func (t *Tree[kT, vT]) Delete(key kT) error {
	del := t.getNode(key)
	if del == t.max {
		t.max = t.max.p
	}
	if del == t.min {
		t.min = t.min.p
	}
	if del == nil {
		return errors.New("key doesn't exist")
	}

	if del.l != nil && del.r != nil {
		next := del.r
		for next.l != nil {
			next = next.l
		}
		del.key = next.key
		del.val = next.val
		del = next
	}

	if del == t.root {
		switch {
		case del.l != nil:
			x := del.l
			x.p = nil
			t.root = x
			x.black = true
		case del.r != nil:
			x := del.r
			x.p = nil
			t.root = x
			x.black = true
		default:
			t.root = nil
		}
		return nil
	}

	var x *node[kT, vT]
	par := del.p
	if del.l != nil {
		x = del.l
		x.p = par
	} else if del.r != nil {
		x = del.r
		x.p = par
	}

	if del == del.p.l {
		par.l = x
	} else {
		par.r = x
	}

	if del.isBlack() {
		for x != t.root && x.isBlack() {
			if x == par.l {
				sib := par.r
				if !sib.isBlack() {
					sib.black = true
					par.black = false
					t.leftRotate(par)
					sib = par.r
				}
				if sib.l.isBlack() && sib.r.isBlack() {
					sib.black = false
					x = par
					par = x.p
				} else {
					if sib.r.isBlack() {
						sib.l.black = true
						sib.black = false
						t.rightRotate(sib)
					}
					sib = par.r
					sib.black = par.black
					par.black = true
					sib.r.black = true
					t.leftRotate(par)
					x = t.root
				}
			} else {
				sib := par.l
				if !sib.isBlack() {
					sib.black = true
					par.black = false
					t.rightRotate(par)
					sib = par.l
				}
				if sib.l.isBlack() && sib.r.isBlack() {
					sib.black = false
					x = par
					par = x.p
				} else {
					if sib.l.isBlack() {
						sib.r.black = true
						sib.black = false
						t.leftRotate(sib)
					}
					sib = par.l
					sib.black = par.black
					par.black = true
					sib.l.black = true
					t.rightRotate(par)
					x = t.root
				}
			}
		}
		t.root.black = true
		x.black = true
	}
	t.size--
	return nil
}

func keyToStr(key any) string {
	switch k := key.(type) {
	case float64:
		s := strconv.FormatFloat(k, 'f', 3, 64)
		return s
	case int:
		s := strconv.Itoa(k)
		return s
	case string:
		return k
	default:
		return ""
	}
}

func (t *Tree[kT, vT]) visualizeInternal(n *node[kT, vT], lBuf string, buf string, rBuf string, kLen int) {
	if n == nil {
		return
	}
	bufLine := rBuf + strings.Repeat(" ", kLen-1) + "│"
	bufSpace := rBuf + strings.Repeat(" ", kLen)
	bufNode := rBuf + strings.Repeat(" ", kLen-1) + "┌"
	t.visualizeInternal(n.r, bufLine, bufNode, bufSpace, kLen)

	var s string
	if n.isBlack() {
		s = keyToStr(n.key) + "(B)"
	} else {
		s = keyToStr(n.key) + "(R)"
	}
	if len(s) < kLen {
		s = strings.Repeat("─", kLen-len(s)) + s
	}
	fmt.Println(buf + s)

	bufLine = lBuf + strings.Repeat(" ", kLen-1) + "│"
	bufSpace = lBuf + strings.Repeat(" ", kLen)
	bufNode = lBuf + strings.Repeat(" ", kLen-1) + "└"
	t.visualizeInternal(n.l, bufSpace, bufNode, bufLine, kLen)
}

// Prints out a visualization of the tree.
// Nodes are represented by their key and color.
// O(n)
func (t *Tree[kT, vT]) Visualize() {
	if t.root == nil {
		fmt.Println()
		return
	}
	kLen := 0
	for _, pair := range t.InOrder() {
		kLen = max(kLen, len(keyToStr(pair[0])))
	}
	kLen += 3
	t.visualizeInternal(t.root, "", "", "", kLen)
}

// Returns the amount of black nodes
// on the path from the root to the leaves.
// O(log(n))
func (t *Tree[kT, vT]) BlackDepth() int {
	n := t.root
	blackDepth := 0
	for n != nil {
		if n.isBlack() {
			blackDepth++
		}
		n = n.r
	}
	return blackDepth
}

func (t *Tree[kT, vT]) verifyInternal(n *node[kT, vT], blackCount int, blackDepth int) error {
	if !n.isBlack() && !n.p.isBlack() {
		return errors.New("red node has a red parent")
	}
	if n.isBlack() {
		blackCount++
	}
	if n.l == nil && n.r == nil && blackCount != blackDepth {
		return errors.New("black depth is inconsistent")
	}
	if n.l != nil {
		err := t.verifyInternal(n.l, blackCount, blackDepth)
		if err != nil {
			return err
		}
	}
	if n.r != nil {
		err := t.verifyInternal(n.r, blackCount, blackDepth)
		if err != nil {
			return err
		}
	}
	return nil
}

// Checks if all the properties of the red-black tree are met.
// Traverses the tree in order and returns an error
// if there is a problem in the tree structure.
// O(n)
func (t *Tree[kT, vT]) Verify() error {
	if t.root == nil {
		return nil
	}
	if !t.root.isBlack() {
		return errors.New("root is not black")
	}
	blackDepth := t.BlackDepth()
	err := t.verifyInternal(t.root, 0, blackDepth)
	return err
}

func (t *Tree[kT, vT]) rangeInternal(n *node[kT, vT], min kT, max kT, out [][]any) [][]any {
	if n.l != nil && n.key >= min {
		out = t.rangeInternal(n.l, min, max, out)
	}
	if n.key >= min && n.key <= max {
		out = append(out, []any{n.key, n.val})
	}
	if n.r != nil && n.key <= max {
		out = t.rangeInternal(n.r, min, max, out)
	}
	return out
}

// Returns a slice of elements not less than min
// and not greater than max.
// O(k), where k is number of such elements.
func (t *Tree[kT, vT]) Range(min kT, max kT) [][]any {
	var out [][]any
	out = t.rangeInternal(t.root, min, max, out)
	return out
}
