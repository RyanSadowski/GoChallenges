package main

import (
	"fmt"
)

func main() {
	var first = isHappy(69)
	var second = isHappy(420)

	fmt.Printf("first: %v, second: %v\n", first, second)
}

// reversed
// pow calculates the power of the given digit
func pow(digit int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * digit
	}
	return res
}

// sumOfSquaredDigits is a helper function that calculates the sum of squared digits
func sumOfSquaredDigits(number int) int {
	totalSum := 0
	for number > 0 {
		digit := number % 10
		number = number / 10
		totalSum += pow(digit, 2)
	}
	return totalSum
}

func isHappy(num int) bool {

	var fast = num
	var slow = num
	for fast != 1 {
		slow = sumOfSquaredDigits(slow)
		fast = sumOfSquaredDigits(fast)
		fast = sumOfSquaredDigits(fast)
		if fast == slow {
			return false
		}
	}

	return true
}
