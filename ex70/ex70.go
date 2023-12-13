/*
	Implements a function that receives a number and return if it is a palindrome
*/

package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	numberInString := strconv.FormatInt(int64(x), 10)

	left := 0
	right := len(numberInString) - 1

	for left < right {
		if numberInString[left] != numberInString[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(-111))
	fmt.Println(isPalindrome(100))
	fmt.Println(isPalindrome(45))
	fmt.Println(isPalindrome(0))
}
