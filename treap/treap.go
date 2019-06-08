package main

import (
	"fmt"
	"math/rand"
)

type Comparable interface {
	less(Comparable) bool
}

type elem_t Comparable
type count_t uint
type pri_t float64
func generateRandomPriority() pri_t {
	return pri_t(rand.Float64())
}

type Node struct {
	elem     elem_t
	lch, rch *Node
	cnt      count_t
	pri      pri_t
}

func NewNode(elem elem_t) *Node {
	node := new(Node)
	node.elem = elem
	node.pri = generateRandomPriority()
	node.cnt = 1
	return node
}

func (node *Node) count() count_t {
	if node != nil {
		return node.cnt
	} else {
		return 0
	}
}

func (node *Node) update() *Node {
	node.cnt = node.lch.count() + node.rch.count() + 1;
	return node
}

func (l *Node) merge(r *Node) *Node {
	// fmt.Println("Node::merge is called")
	// fmt.Println("Node::merge l:",l,", r:",r)
	// fmt.Println()
	if l == nil {
		return r
	} else if r == nil {
		return l
	}

	l.update()
	r.update()
	if (l.pri > r.pri) {
		l.rch = l.rch.merge(r)
		return l.update()
	} else {
		r.lch = l.merge(r.lch)
		return r.update()
	}
}

func (node *Node) split(elem elem_t, includeToRight bool) (*Node, *Node) {
	// fmt.Println("Node::split is called")
	// fmt.Println("Node::split node:",node,", elem:",elem)
	// fmt.Println()
	if node == nil {
		return nil, nil
	}
	node.update()

	if includeToRight {
		if node.elem.less(elem) {
			l, r := node.rch.split(elem, includeToRight)
			node.rch = l
			return node.update(), r
		} else {
			l, r := node.lch.split(elem, includeToRight)
			node.lch = r
			return l, node.update()
		}
	} else {
		if elem.less(node.elem) {
			l, r := node.lch.split(elem, includeToRight)
			node.lch = r
			return l, node.update()
		} else {
			l, r := node.rch.split(elem, includeToRight)
			node.rch = l
			return node.update(), r
		}
	}
}

func (node *Node) insert(elem elem_t) *Node {
	// fmt.Println("Node::insert is called")
	// fmt.Println("Node::insert elem:", elem)
	// fmt.Println()
	c, r := node.split(elem, false)
	l, _ := c.split(elem, true)
	// fmt.Println("Node::insert l:",l,", r:",r)
	return l.merge(NewNode(elem).merge(r))
}

func (node *Node) erase(elem elem_t) (*Node, *Node) {
	c, r := node.split(elem, false)
	l, _ := c.split(elem, true)
	return l.merge(r), c
}

func (node *Node) find(elem elem_t) *Node {
	if node == nil {
		return nil
	}
	if node.elem.less(elem) {
		return node.rch.find(elem)
	} else if elem.less(node.elem) {
		return node.lch.find(elem)
	} else {
		return node
	}
}

func (node *Node) show() {
	if node == nil {
		return
	}
	node.lch.show()
	fmt.Println("Node::show node:", node)
	node.rch.show()
}


type Treap struct {
	root *Node
}

func (treap *Treap) size() count_t {
	return treap.root.count()
}

func (treap *Treap) insert(elem elem_t) {
	fmt.Println("Treap::insert is called")
	treap.root = treap.root.insert(elem)
}

func (treap *Treap) erase(elem elem_t) {
	fmt.Println("Treap::erase is called")
	treap.root, _ = treap.root.erase(elem)
}

func (treap *Treap) find(elem elem_t) *Node {
	fmt.Println("Treap::find is called")
	return treap.root.find(elem)
}

func (treap *Treap) show() {
	treap.root.show()
}

type Int int

func (item Int) less(other Comparable) bool {
	otherT, ok := other.(Int)
	if !ok {
		//handle error (other was not of type T)
	}
	return item < otherT
}

func main() {
	rand.Seed(2)
	treap := new(Treap)

	insertList := []int{3, 3, 5, 4, 2, 4}
	for i, v := range insertList {
		fmt.Println("i:",i)
		treap.insert(Int(v))
		treap.show()
	}
	findList := []int{3, 5, 1}
	for i, v := range findList {
		fmt.Println("i:",i)
		fmt.Println("treap.find(",v,"):",treap.find(Int(v)))
	}
	eraseList := []int{3, 3, 2, 4, 5}
	for i, v := range eraseList {
		fmt.Println("i:",i)
		treap.erase(Int(v))
		treap.show()
	}
}
