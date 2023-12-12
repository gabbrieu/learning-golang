package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func (t *Tree) search(num int) bool {
	currentNode := t.root
	return t.searchRecursive(num, currentNode)

}

func (t *Tree) searchRecursive(num int, currentNode *Node) bool {
	if num == currentNode.value {
		return true
	}

	if num < currentNode.value {
		if currentNode.left == nil {
			return false
		}
		return t.searchRecursive(num, currentNode.left)
	}

	if currentNode.right == nil {
		return false
	}
	return t.searchRecursive(num, currentNode.right)
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

	fmt.Print("Tree in order: ")
	binaryTree.printInOrder()

	fmt.Print("\n\nEnter a number to search on the Tree: ")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')

	userInputConverted, err := strconv.Atoi(strings.Trim(userInput, " \n"))

	if err != nil {
		panic(err)
	}

	numberExist := binaryTree.search(userInputConverted)

	if numberExist {
		fmt.Printf("The number %d exists on the Tree", userInputConverted)
	} else {
		fmt.Printf("The number %d does NOT exist on the Tree", userInputConverted)
	}
}
