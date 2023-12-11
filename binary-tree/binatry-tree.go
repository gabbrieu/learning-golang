package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	value int
}

type Tree struct {
	root *Node
}

func (t *Tree) add(num int) {
	newNode := &Node{value: num}

	if t.root == nil {
		t.root = newNode
		return
	}

	currentNode := t.root

	t.moveToAdd(num, currentNode, newNode)
}

func (t *Tree) moveToAdd(num int, currentNode *Node, newNode *Node) {
	if num > currentNode.value {
		if currentNode.right == nil {
			currentNode.right = newNode
			return
		}

		t.moveToAdd(num, currentNode.right, newNode)

	} else {
		if currentNode.left == nil {
			currentNode.left = newNode
			return
		}

		t.moveToAdd(num, currentNode.left, newNode)
	}
}

func (t *Tree) printInOrder() {
	if t.root == nil {
		fmt.Print("Empty tree")
		return
	}

	currentNode := t.root

	t.inOrder(currentNode)
}

func (t *Tree) inOrder(currentNode *Node) {
	if currentNode == nil {
		return
	}

	if currentNode.left == nil && currentNode.right == nil {
		fmt.Printf("%d ", currentNode.value)
		return
	}

	if currentNode.left != nil {
		t.inOrder(currentNode.left)
	}

	fmt.Printf("%d ", currentNode.value)
	t.inOrder(currentNode.right)
}

func main() {
	binaryTree := &Tree{}

	binaryTree.add(100)
	binaryTree.add(-20)
	binaryTree.add(-50)
	binaryTree.add(-15)
	binaryTree.add(-60)
	binaryTree.add(50)
	binaryTree.add(60)
	binaryTree.add(55)
	binaryTree.add(85)
	binaryTree.add(15)
	binaryTree.add(5)
	binaryTree.add(-10)

	binaryTree.printInOrder()
}
