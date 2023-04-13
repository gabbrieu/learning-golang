/**
- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
**/

package main

import "fmt"

func main() {
	sliceOfInt := []int{1, 2, 3, 4}
	fmt.Println(sum(sliceOfInt...))
	fmt.Println(sumSlice(sliceOfInt))
}

func sum(x ...int) int {
	var soma int = 0

	for _, v := range x {
		soma += v
	}
	return soma
}

func sumSlice(x []int) int {
	var soma int = 0

	for _, v := range x {
		soma += v
	}
	return soma
}
