package main

import (
	"fmt"
)

type Node struct {
	next  *Node
	value int
}

type List struct {
	head *Node
}

func (l *List) add(num int) {
	n := &Node{value: num}

	if l.head == nil {
		n.next = nil
		l.head = n
		return
	}

	currentNode := l.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = n
}

func (l *List) remove(num int) {
	if l.head == nil {
		return
	}

	currentNode := l.head

	// Removing first node
	if l.head.value == num {
		l.head = currentNode.next
		currentNode.next = nil
		return
	}

	for currentNode.next != nil && currentNode.next.value != num {
		currentNode = currentNode.next
	}

	currentNode.next = currentNode.next.next
}

func (l *List) pop() {
	if l.head == nil {
		return
	}

	if l.head.next == nil {
		l.head = nil
		return
	}

	currentNode := l.head
	previousNode := l.head

	for currentNode.next != nil {
		previousNode = currentNode
		currentNode = currentNode.next
	}

	previousNode.next = nil

}

func (l *List) print() {
	currentNode := l.head

	for currentNode != nil {
		fmt.Printf("%d ", currentNode.value)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func main() {
	linkedList := &List{}
	linkedList.add(10)
	linkedList.add(22)
	linkedList.add(26)
	linkedList.add(34)
	linkedList.add(50)

	fmt.Print("Full list: ")
	linkedList.print()

	linkedList.remove(10)
	fmt.Print("After remove 10: ")
	linkedList.print()

	linkedList.remove(22)
	fmt.Print("After remove 22: ")
	linkedList.print()

	linkedList.remove(50)
	fmt.Print("After remove 50: ")
	linkedList.print()

	linkedList.pop()
	fmt.Print("After pop: ")
	linkedList.print()

	linkedList.pop()
	fmt.Print("After pop: ")
	linkedList.print()
}
