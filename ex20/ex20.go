// Tente acessar todos os itens de uma slice sem utilizar range.
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}

	fmt.Println(slice[:])
}
