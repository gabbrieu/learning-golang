/**
- Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
**/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func1()
	go func2()

	wg.Wait()
}

func func1() {
	fmt.Println("I'm the first function")
	wg.Done()
}

func func2() {
	fmt.Println("I'm the second function")
	wg.Done()
}
