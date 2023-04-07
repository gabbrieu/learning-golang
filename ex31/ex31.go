/**
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
**/

package main

import "fmt"

func main() {
	type pessoa struct {
		nome                    string
		sobrenome               string
		saboresFavoritosSorvete []string
	}

	pessoa1 := pessoa{
		nome:                    "Gabriel",
		sobrenome:               "Mendes",
		saboresFavoritosSorvete: []string{"Chocomenta", "Brigadeiro"},
	}

	pessoa2 := pessoa{"Julia", "Maria", []string{"Morango", "Creme"}}

	fmt.Println("PESSOA 1")
	fmt.Println(pessoa1.nome, pessoa1.sobrenome)

	for _, v := range pessoa1.saboresFavoritosSorvete {
		fmt.Printf("\t%v", v)
	}

	fmt.Println("\n\nPESSOA 2")
	fmt.Println(pessoa2.nome, pessoa2.sobrenome)

	for _, v := range pessoa2.saboresFavoritosSorvete {
		fmt.Printf("\t%v", v)
	}
}
