// Construa a função fatorial usando recursividade

package main

import "fmt"

func main() {
	fmt.Println(fatorial(9))
}

func fatorial(x int) int {
	if x == 1 {
		return 1
	}

	return fatorial(x-1) * x
}
