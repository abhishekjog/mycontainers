package mylist

import (
	"mycontainers/node"
)

type SingleLinkedList struct {
	Head *node.Node
}

func (l *SingleLinkedList) Push(data int) *node.Node {
	temp := &node.Node{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = temp
	} else {
		var trav *node.Node
		for trav = l.Head; trav.Next != nil; trav = trav.Next {
		}
		trav.Next = temp
	}
	return l.Head
}

func (l *SingleLinkedList) Pop() (int, bool) {
	ret := l.Head
	if ret == nil {
		return 0, false
	} else {
		l.Head = l.Head.Next
		return ret.Data, true
	}
}
