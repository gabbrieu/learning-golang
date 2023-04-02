/**
- Crie constantes tipadas e não-tipadas.
- Demonstre seus valores.
**/

package main

import "fmt"

func main() {
	const x int = 10
	const y = "Olá"

	fmt.Printf("x: %v\tTipo de x: %T\ny: %v\tTipo de y: %T\n", x, x, y, y)
}
