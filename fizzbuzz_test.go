package fizzbuzz_test

import (
	. "fizzbuzz"
	"testing"
)

func TestGivenFizzBuzzSay30ShouldBeFizzBuzz(t *testing.T) {
	Given{
		t,
		FizzBuzz{}.Say(30),
	}.ShouldBe("FizzBuzz")
}

func TestGivenFizzBuzzSay15ShouldBeFizzBuzz(t *testing.T) {
	Given{
		t,
		FizzBuzz{}.Say(15),
	}.ShouldBe("FizzBuzz")
}

func TestGivenFizzBuzzSay10ShouldBeBuzz(t *testing.T) {
	Given{
		t,
		FizzBuzz{}.Say(10),
	}.ShouldBe("Buzz")
}

func TestGivenFizzBuzzSay6ShouldBeFizz(t *testing.T) {
	Given{
		t,
		FizzBuzz{}.Say(6),
	}.ShouldBe("Fizz")
}

func TestGivenFizzBuzzSay5ShouldBeBuzz(t *testing.T) {
	Given{
		t,
		FizzBuzz{}.Say(5),
	}.ShouldBe("Buzz")
}

func TestGivenFizzBuzzSay3ShouldBeFizz(t *testing.T) {
	Given{
		t,
		FizzBuzz{}.Say(3),
	}.ShouldBe("Fizz")
}

func TestGivenFizzBuzzSay2ShouldBe2(t *testing.T) {
	Given{
		t,
		FizzBuzz{}.Say(2),
	}.ShouldBe("2")
}

func TestGivenFizzBuzzSay1ShouldBe1(t *testing.T) {
	Given{
		t,
		FizzBuzz{}.Say(1),
	}.ShouldBe("1")
}
