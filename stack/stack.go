package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) push(num int) {
	newNode := &Node{value: num}

	if s.top == nil {
		s.top = newNode
		return
	}

	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) pop() {
	if s.top == nil {
		return
	}

	temp := s.top.next
	s.top.next = nil
	s.top = temp
}

func (s *Stack) print() {
	if s.top == nil {
		return
	}

	currentNode := s.top

	for currentNode != nil {
		fmt.Printf("%d ", currentNode.value)
		currentNode = currentNode.next
	}
}

func main() {
	stack := &Stack{}

	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)

	stack.print()

	stack.pop()

	fmt.Print("\nAfter pop: ")
	stack.print()

}
