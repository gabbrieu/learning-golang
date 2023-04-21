/**
- Crie um programa que demonstra seu OS e ARCH.
- Rode-o com os seguintes comandos:
    - go run
    - go build
    - go install
**/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Meu OS é:", runtime.GOOS)
	fmt.Println("Minha ARCH é:", runtime.GOARCH)
}
