package main

import (
	"fmt"
)

func isValidPalindrome(pp string) bool {
	var left int = 0
	var right int = len(pp) - 1
	var hasSkipped bool = false

	for left < right {
		if pp[left] != pp[right] {
			if hasSkipped {
				return false
			}

			//peek
			if pp[left] == pp[right-1] {
				right--
			} else if pp[right] == pp[left+1] {
				left++
			} else {
				return false
			}

			hasSkipped = true
		}
		left++
		right--
	}

	return true
}

func main() {
	var pp1 string = "tacocat"  //true
	var pp2 string = "tacomcat" //true, remove the m
	var pp3 string = "tacoma"   //false

	fmt.Println(pp1)
	fmt.Println(isValidPalindrome(pp1))
	fmt.Println(pp2)
	fmt.Println(isValidPalindrome(pp2))
	fmt.Println(pp3)
	fmt.Println(isValidPalindrome(pp3))

}
