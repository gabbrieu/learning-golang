/**
- Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados
**/

package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	years, name := info()
	fmt.Printf("Hello, my name is %v and i have %v years old", name, years)
}

func sum(x ...int) int {
	var sum int = 0

	for _, v := range x {
		sum += v
	}
	return sum
}

func info() (int, string) {
	return 23, "Gabriel"
}
