package linkedlist

import "errors"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.

type List struct {
	first *Node
	last  *Node
}

type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

func NewList(args ...interface{}) *List {
	var l List
	for _, n := range args {
		l.Push(n)
	}

	return &l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	newNode := &Node{Value: v}

	if l.first == nil {
		l.first = newNode
		l.last = newNode
		return
	}

	oldFirst := l.first

	oldFirst.prev = newNode
	newNode.next = oldFirst
	l.first = newNode
}

func (l *List) Push(v interface{}) {
	newNode := &Node{Value: v}

	if l.first == nil {
		l.first = newNode
		l.last = newNode
		return
	}

	oldLast := l.last

	oldLast.next = newNode
	newNode.prev = oldLast
	l.last = newNode
}

func (l *List) Shift() (interface{}, error) {
	if l.first == nil {
		return nil, errors.New("empty list")
	}

	shifted := l.first
	newFirst := l.first.next

	if newFirst == nil {
		l.first = nil
		l.last = nil
		return shifted.Value, nil
	}

	newFirst.prev = nil
	l.first = newFirst

	return shifted.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.last == nil {
		return nil, errors.New("empty list")
	}

	poped := l.last
	newLast := l.last.prev

	if newLast == nil {
		l.first = nil
		l.last = nil
		return poped.Value, nil
	}

	l.last = newLast
	newLast.next = nil

	return poped.Value, nil
}

func (l *List) Reverse() {
	if l.last == nil {
		return
	}

	l.first, l.last = l.last, l.first

	n := l.first

	for n != nil {
		n.prev, n.next = n.next, n.prev
		n = n.next
	}
}

//func (l *List) Reverse() {
//	if l.last == nil {
//		return
//	}
//
//	revertNode(l.last)
//
//	l.first, l.last = l.last, l.first
//}

func revertNode(n *Node) {
	if n.prev == nil {
		n.prev, n.next = n.next, nil
		return
	}

	revertNode(n.prev)

	n.next, n.prev = n.prev, n.next
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
