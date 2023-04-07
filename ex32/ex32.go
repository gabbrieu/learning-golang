/**
- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
**/

package main

import "fmt"

func main() {
	type pessoa struct {
		nome                    string
		sobrenome               string
		saboresFavoritosSorvete []string
	}

	// ------------ Other way to achieve that!
	// pessoa1 := pessoa{
	// 	nome:                    "Gabriel",
	// 	sobrenome:               "Mendes",
	// 	saboresFavoritosSorvete: []string{"Chocomenta", "Brigadeiro"},
	// }

	// x := map[string]pessoa{
	// 	pessoa1.sobrenome: pessoa1,
	// }

	myMap := make(map[string]pessoa)

	myMap["Mendes"] = pessoa{
		nome:                    "Gabriel",
		sobrenome:               "Mendes",
		saboresFavoritosSorvete: []string{"Chocomenta", "Brigadeiro"},
	}

	for _, value := range myMap {
		fmt.Printf("Sou %v %v", value.nome, value.sobrenome)

		fmt.Print(" e gosto de sorvete de: \n")
		for _, v := range value.saboresFavoritosSorvete {
			fmt.Printf("- %v \n", v)
		}
	}
}
