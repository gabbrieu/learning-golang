/**
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
**/

package main

import "fmt"

func main() {
	ano := 1999
	for ano <= 2023 {
		fmt.Println(ano)
		ano++
	}
}