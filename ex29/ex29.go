// - Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.

package main

import "fmt"

func main() {
	x := map[string][]string{
		"mendes_gabriel": {"jogar", "estudar", "sair"},
		"julia_maria":    {"fazer ovo", "me irritar"},
	}

	x["einstein_albert"] = []string{"estudar", "ganhar nobels"}

	for i, v := range x {
		fmt.Println(i)

		for j, v1 := range v {
			fmt.Printf("\t%v - %v\n", j, v1)
		}
	}
}
