// - Crie um programa que utilize a declaração switch, sem switch statement especificado.
package main

import "fmt"

func main() {
	itsWeekend := false

	switch {
	case itsWeekend == true:
		fmt.Println("Uhu! É Fim de semana")

	case itsWeekend == false:
		fmt.Println("Ahh! Tenho que trabalhar :(")
	}
}
