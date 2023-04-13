// - Crie e utilize uma função anônima.

package main

import "fmt"

func main() {
	x, y := 4, 7

	func(x int, y int) {
		fmt.Println(x + y)
	}(x, y)
}
