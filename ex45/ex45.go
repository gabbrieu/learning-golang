//- Callback: passe uma função como argumento a outra função.

package main

import "fmt"

func main() {
	anotherFunction(callback)
}

func callback() {
	fmt.Println("Are you there?")
}

func anotherFunction(x func()) {
	fmt.Println("Hey:")
	x()
}
