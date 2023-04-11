/*
*
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*
*/
package main

import "fmt"

func main() {
	myType := struct {
		myMap   map[string]string
		name    string
		mySlice []string
	}{
		name:    "myType",
		myMap:   map[string]string{"test": "test"},
		mySlice: []string{"first element", "second element"},
	}

	fmt.Println(myType)
}
