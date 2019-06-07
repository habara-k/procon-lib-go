package main

import (
	"fmt"
)


type Comparable interface {
	less(Comparable) bool
}


type Node struct {
	elem Comparable
	par, left, right *Node
}

func NewNode(elem Comparable, par *Node) *Node {
    node := new(Node)
	node.elem = elem
	node.par = par
    return node
}

func (node *Node) insert(elem Comparable, par *Node) *Node {
	// fmt.Println("Node::insert is called")
	if node == nil {
		node = NewNode(elem, par)
		if par == nil { return node }
		if par.elem.less(elem) {
			par.right = node
		} else {
			par.left = node
		}
		return node
	}
	if node.elem.less(elem) {
		node.right.insert(elem, node)
	} else if elem.less(node.elem) {
		node.left.insert(elem, node)
	}
	return node
}

func (node *Node) find(elem Comparable) *Node {
	// fmt.Println("Node::find is called")
	if node == nil {
		return nil
	}
	if node.elem.less(elem) {
		return node.right.find(elem)
	} else if elem.less(node.elem) {
		return node.left.find(elem)
	} else {
		return node
	}
}

func (node *Node) getMin() *Node {
	// fmt.Println("Node::getMin is called")
	if node.left == nil {
		return node
	} else {
		return node.left.getMin()
	}
}

func (node *Node) getMax() *Node {
	// fmt.Println("Node::getMax is called")
	if node.right == nil {
		return node
	} else {
		return node.right.getMax()
	}
}

func (node *Node) replace() {
	// fmt.Println("Node::replace is called")
	var child *Node
	if node.left == nil {
		child = node.right
	} else {
		child = node.left
	}

	if child != nil {
		child.par = node.par
	}

	if node.par != nil {
		if node.par.elem.less(node.elem) {
			node.par.right = child
		} else {
			node.par.left = child
		}
	}
}

func (node *Node) erase() {
	// fmt.Println("Node::erase is called")
	if node == nil {
		return
	}

	if node.left != nil && node.right != nil {
		rightMinNode := node.right.getMin()
		node.elem, rightMinNode.elem = rightMinNode.elem, node.elem
		rightMinNode.erase()
	} else {
		node.replace()
	}
}

func (node *Node) next() *Node {
	if node.right != nil {
		return node.right.getMin()
	} else {
		n := node
		for {
			if n.par == nil { return nil }
			if n.elem.less(n.par.elem) { return n.par }
			n = n.par
		}
	}
}

func (node *Node) show() {
	if node == nil {
		return
	}
	node.left.show()
	fmt.Println(node.elem)
	node.right.show()
}


type Set struct {
	root *Node
}

func (set *Set) insert(elem Comparable) {
	fmt.Println("Set::insert is called")
	set.root = set.root.insert(elem, nil)
}

func (set *Set) find(elem Comparable) *Node {
	fmt.Println("Set::find is called")
	return set.root.find(elem)
}

func (set *Set) erase(elem Comparable) {
	fmt.Println("Set::erase is called")
	node := set.root.find(elem)
	if node == set.root {
		set.root = nil
	} else {
		node.erase()
	}
}

func (set *Set) begin() *Node {
	return set.root.getMin()
}

func (set *Set) show() {
	set.root.show()
}


type Int int
func (item Int) less(other Comparable) bool {
	otherT, ok := other.(Int)
    if !ok{
        //handle error (other was not of type T)
    }
	return item < otherT
}


func main() {
	set := new(Set)

	set.insert(Int(3))
	set.show()

	set.erase(Int(3))
	set.show()

	for i := 0; i < 8; i++ {
		set.insert(Int(i))
	}
	set.show()

	fmt.Println(set.find(Int(2)))
	fmt.Println(set.find(Int(10)))

	for it := set.begin(); it != nil; it = it.next() {
		fmt.Println(it.elem)
	}
}

