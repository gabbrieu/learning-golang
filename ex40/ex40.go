/**
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
**/

package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      int
}

func (p person) showPersonInfo() {
	fmt.Printf("My full name is: %v %v\nMy age is: %v", p.name, p.lastName, p.age)
}

func main() {
	var gabriel person = person{
		name:     "Gabriel",
		lastName: "Mendes",
		age:      23,
	}

	gabriel.showPersonInfo()
}
