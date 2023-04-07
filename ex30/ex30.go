// - Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.

package main

import "fmt"

func main() {
	x := map[string][]string{
		"mendes_gabriel": {"jogar", "estudar", "sair"},
		"maria_julia":    {"fazer ovo", "me irritar"},
	}

	x["einstein_albert"] = []string{"estudar", "ganhar nobels"}

	delete(x, "maria_julia")

	for i, v := range x {
		fmt.Println(i)

		for j, v1 := range v {
			fmt.Printf("\t%v - %v\n", j, v1)
		}
	}
}
