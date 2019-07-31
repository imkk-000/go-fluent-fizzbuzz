package fizzbuzz

import (
	"strconv"
)

// FizzBuzz - this is the group of FizzBuzz methods
type FizzBuzz struct{}

var divisorOrderTable = []int{15, 3, 5}
var ruleTable = map[int]string{
	15: "FizzBuzz",
	3:  "Fizz",
	5:  "Buzz",
}

// Say - when you play the game
// you must say "Fizz", "Buzz", "FizzBuzz", or natual number
func (FizzBuzz) Say(number int) string {
	for _, divisor := range divisorOrderTable {
		if number%divisor == 0 {
			return ruleTable[divisor]
		}
	}
	return strconv.Itoa(number)
}
