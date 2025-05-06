package sample

import "fmt"

func fizzBuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return fmt.Sprintf("%d", n)
	}
}

// Added a new feature to handle negative numbers
func fizzBuzzWithNegative(n int) string {
	if n < 0 {
		return "Negative numbers are not allowed"
	}
	return fizzBuzz(n)
}
