/**
- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
**/

package main

import "fmt"

func main() {
	x := [][]string{
		{"Gabriel", "Mendes", "Jogar"},
		{"Julia", "Maria", "Me irritar"},
		{"Albert", "Einstein", "Estudar"},
	}

	for i := range x {
		fmt.Println(x[i])
	}
}
