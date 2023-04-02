// - Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		moduleByFour := i % 4
		fmt.Printf("Módulo de %v por 4: %v\n", i, moduleByFour)
	}
}
