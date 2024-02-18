package tree

import (
	"errors"
	"golang.org/x/exp/constraints"
	"strconv"
	"strings"
	"fmt"
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
		s := strconv.FormatFloat(k, 'f', 2, 64)
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
	bufLine := rBuf + strings.Repeat(" ", kLen - 1) + "│"
	bufSpace := rBuf + strings.Repeat(" ", kLen)
	bufNode := rBuf + strings.Repeat(" ", kLen - 1) + "┌"
	t.visualizeInternal(n.r, bufLine, bufNode, bufSpace,  kLen)

	var s string
	if n.isBlack() {
		s = keyToStr(n.key) + "(B)"
	} else {
		s = keyToStr(n.key) + "(R)"
	}
	if len(s) < kLen {
		s = strings.Repeat("─", kLen - len(s)) + s
	}
	fmt.Println(buf + s)

	bufLine = lBuf + strings.Repeat(" ", kLen - 1) + "│"
	bufSpace = lBuf + strings.Repeat(" ", kLen)
	bufNode = lBuf + strings.Repeat(" ", kLen - 1) + "└"
	t.visualizeInternal(n.l, bufSpace, bufNode, bufLine,  kLen)
}

func (t *Tree[kT, vT]) Visualize(){
	kLen := len(keyToStr(t.max.key)) + 3
	t.visualizeInternal(t.root, "", "", "", kLen)
}
