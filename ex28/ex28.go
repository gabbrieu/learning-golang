/**
- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
**/

package main

import "fmt"

func main() {
	x := map[string][]string{
		"mendes_gabriel": {"jogar", "estudar", "sair"},
		"julia_maria":    {"fazer ovo", "me irritar"},
	}

	for i, v := range x {
		fmt.Println(i)

		for j, v1 := range v {
			fmt.Printf("\t%v - %v\n", j, v1)
		}
	}
}
