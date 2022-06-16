package main

import (
	"fmt"
)

type Node struct {
	next *Node
	key  int
}

type List struct {
	head     *Node
	nextHead *List
	tail     *Node
}

func merge(L1 *List, L2 *List) *List {
	l1 := L1.head
	l2 := L2.head
	var dummy = new(Node)
	var p = dummy
	for l1 != nil && l2 != nil {
		if l1.key < l2.key {
			p.next = l1
			l1 = l1.next
		} else {
			p.next = l2
			l2 = l2.next
		}
		p = p.next
	}
	if l1 != nil {
		p.next = l1
	} else {
		p.next = l2
	}
	link := List{}
	link.head = dummy.next
	return &link

}
func (L *List) Insert(key int) {
	list := &Node{
		next: nil,
		key:  key,
	}

	if L.head != nil {
		L.tail.next = list
	} else {
		L.head = list

	}
	L.tail = list
}

func (node *Node) display() {
	for {
		if node == nil {
			break
		}
		fmt.Println(node.key)
		node = node.next
	}
}

func _merge(l *List) *List {
	// tmp:=l
	tmp1 := l.nextHead
	for tmp1 != nil {
		res := merge(l, tmp1)
		l = res
		tmp1 = tmp1.nextHead

	}
	return l
}

func main() {
	link1 := List{}
	link1.Insert(5)
	link1.Insert(9)
	link1.Insert(13)
	link1.Insert(22)
	link1.Insert(28)
	link1.Insert(44)

	link2 := List{}

	link2.Insert(6)
	link2.Insert(7)
	link2.Insert(20)
	link2.Insert(25)
	link2.Insert(40)
	link2.Insert(45)
	link1.nextHead = &link2

	link3 := List{}
	link3.Insert(61)
	link3.Insert(17)
	link3.Insert(31)
	link3.Insert(32)
	link3.Insert(41)
	link3.Insert(49)
	link2.nextHead = &link3

	res := _merge(&link1)
	res.head.display()
}
