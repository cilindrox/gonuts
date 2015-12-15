// Package gonuts implements the exercises found in An Introduction to Programming in Go.
// See https://www.golang-book.com
package gonuts

import "fmt"

// FizzBuzz returns the string "Fizz" for multiples of 3,
// "Buzz" for multiples of 5 and "FizzBuzz" for numbers that satisfy both conditions.
// If none of the above apply, the current number, formatted as a string, will be returned.
func FizzBuzz(i int) string {

	result := ""
	if i%3 == 0 || i%5 == 0 {
		if i%3 == 0 {
			result += fmt.Sprintf("Fizz")
		}
		if i%5 == 0 {
			result += fmt.Sprintf("Buzz")
		}
	} else {
		result += fmt.Sprintf("%d", i)
	}

	return result
}
