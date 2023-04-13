/**
- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.
**/

package main

import "fmt"

func main() {
	a := returnsAFunction()

	fmt.Println(a(3, 5))
}

func returnsAFunction() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}
