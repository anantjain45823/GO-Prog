package main

import (
	"fmt"
)

type Node struct {
	next *Node
	key  int
}

type List struct {
	head *Node
	tail *Node
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

func ReverseKGroup(head *Node, k int) *Node {

	curr := head
	var prev *Node
	var next *Node
	i := 0

	for curr != nil && i < k {
		i++
		curr = curr.next
	}
	if i == k {
		curr = head
	} else {
		return head
	}

	i = 0
	for curr != nil && i < k {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
		i++
	}

	head.next = ReverseKGroup(curr, k)
	return prev
}

func main() {
	k := 4
	link1 := List{}
	link1.Insert(5)
	link1.Insert(9)
	link1.Insert(13)
	link1.Insert(22)
	link1.Insert(28)
	link1.Insert(44)
	fmt.Print("Before:")
	link1.head.display()
	fmt.Println("After:")
	res := ReverseKGroup(link1.head, k)
	res.display()
}
