/**
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.
**/

package main

import "fmt"

func main() {
	const (
		_ = iota + 2023
		a
		b
		c
		d
	)

	fmt.Println(a, b, c, d)
}
