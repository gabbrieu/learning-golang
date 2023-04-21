/**
- Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
    - Crie um tipo para um struct chamado "pessoa"
    - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
    - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
    - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
    - Demonstre no seu código:
        - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
        - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
**/

package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) talk() {
	fmt.Printf("Hi! I'm %v and i have %v years old", p.name, p.age)
}

type Humans interface {
	talk()
}

func saySomething(h Humans) {
	h.talk()
}

func main() {
	var personOne Person = Person{"Gabriel", 23}

	saySomething(&personOne)
	// saySomething(personOne) cannot use personOne (variable of type Person) as Humans value in argument to saySomething: Person does not implement Humans (method talk has pointer receiver)
}
