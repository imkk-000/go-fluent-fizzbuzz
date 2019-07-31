package fizzbuzz

import (
	"strconv"
)

// FizzBuzz - this is the group of FizzBuzz methods
type FizzBuzz struct{}

// Say - when you play the game
// you must say "Fizz", "Buzz", "FizzBuzz", or natual number
func (FizzBuzz) Say(number int) string {
	switch {
	case number%15 == 0:
		return "FizzBuzz"
	case number%3 == 0:
		return "Fizz"
	case number%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(number)
}
