/**
- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: π * raio * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
**/

package main

import (
	"fmt"
	"math"
)

type square struct {
	side int
}

type circle struct {
	radius int
}

func (s square) area() float64 {
	return float64(s.side * s.side)
}

func (c circle) area() float64 {
	return float64(math.Pi * float64(c.radius) * float64(c.radius))
}

type figure interface {
	area() float64
}

func info(f figure) float64 {
	return f.area()
}

func main() {
	var (
		square square = square{3}
		circle circle = circle{2}
	)

	fmt.Printf("Area of the square: %v\nArea of the circle: %v", info(square), info(circle))
}
