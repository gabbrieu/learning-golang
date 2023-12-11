/** Faça esse código funcionar
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}
**/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
