/** use um select statement para demonstrar os valores do canal.
package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}

	return c
}
**/

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		close(c)

		q <- 1
	}()

	return c
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case v, ok := <-c:
			if ok {
				fmt.Println(v)
			}
		case <-q:
			return
		}
	}
}
