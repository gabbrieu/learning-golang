/**
- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
**/

package main

import "fmt"

func main() {
	ano := 1999
	for {
		fmt.Println(ano)
		ano++
		if ano > 2023 {
			break
		}
	}
}
