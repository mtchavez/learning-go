package main

import (
	"fmt"
)

type Value int

type Node struct {
	Value
	prev *Node
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Front() *Node {
	return l.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (l *List) Push(v Value) *List {
	n := &Node{Value: v}
	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
		n.prev = l.tail
	}
	l.tail = n

	return l
}

func (l *List) Pop() (val Value, err bool) {
	if l.tail == nil {
		err = true
	} else {
		val = l.tail.Value
		l.tail = l.tail.prev
		if l.tail == nil {
			l.head = nil
		}
	}
	return
}

func (l *List) Print() {
	fmt.Println("List Contains:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d\n", e.Value)
	}
	fmt.Println()
}

func main() {
	l := new(List)
	l.Push(8)
	l.Push(9)
	l.Push(10)

	l.Print()

	l.Pop()
	l.Pop()
	l.Pop()
	l.Print()
	_, err := l.Pop()
	if err {
		fmt.Println("error popping from list")
	}
}
