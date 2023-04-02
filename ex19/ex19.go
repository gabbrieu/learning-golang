// - Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
package main

import "fmt"

func main() {
	esporteFavorito := "Futebol"

	switch esporteFavorito {
	case "Futebol":
		fmt.Println("Bora bater uma bolinha!")

	case "Basquete":
		fmt.Println("Um basquetezinho é dentro!")

	default:
		fmt.Println("Sem esporte favorito")
	}
}
