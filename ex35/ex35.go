// - Criando uma função que toma uma função e um []int, e usa somente os números pares como argumentos para a função.

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	evenSum := sumEven(sum, numbers...)
	oddSum := sumOdd(sum, numbers...)

	fmt.Printf("The sum of the even numbers is: %v\nThe sum of the odd numbers is: %v", evenSum, oddSum)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}

func sumEven(f func(x ...int) int, numbers ...int) int {
	var evenSlice []int

	for _, v := range numbers {
		if v%2 == 0 {
			evenSlice = append(evenSlice, v)
		}
	}
	return f(evenSlice...)
}

func sumOdd(f func(x ...int) int, numbers ...int) int {
	var oddSlice []int

	for _, v := range numbers {
		if v%2 != 0 {
			oddSlice = append(oddSlice, v)
		}
	}

	return f(oddSlice...)
}
