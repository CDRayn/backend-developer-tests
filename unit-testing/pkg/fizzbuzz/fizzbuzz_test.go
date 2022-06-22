package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const fizz = "Fizz"
const buzz = "Buzz"
const fizzbuzz = "FizzBuzz"

type TestCase struct {
	total, fizzAt, buzzAt int64
	result                []string
}

// TestFizzBuzzValid tests the behavior of FizzBuzz() using expected values for its params
func TestFizzBuzzValid(t *testing.T) {

	validTestCases := []TestCase{
		{
			total:  10,
			fizzAt: 2,
			buzzAt: 5,
			result: []string{"1", fizz, "3", fizz, buzz, fizz, "7", fizz, "9", fizzbuzz},
		},
		{
			total:  10,
			fizzAt: 5,
			buzzAt: 2,
			result: []string{"1", buzz, "3", buzz, fizz, buzz, "7", buzz, "9", fizzbuzz},
		},
		{
			total:  5,
			fizzAt: 2,
			buzzAt: 5,
			result: []string{"1", fizz, "3", fizz, buzz},
		},
		{
			total:  5,
			fizzAt: 2,
			buzzAt: 2,
			result: []string{"1", fizzbuzz, "3", fizzbuzz, "5"},
		},
		// Test for when a total value is less than fizz or buzz
		{
			total:  2,
			fizzAt: 3,
			buzzAt: 5,
			result: []string{"1", "2"},
		},
	}

	for _, element := range validTestCases {
		assert.Equal(t, FizzBuzz(element.total, element.fizzAt, element.buzzAt), element.result)
	}
}

// TestFizzBuzz_0Total tests the behavior of FizzBuzz() when the total parameter is 0
func TestFizzBuzz_0Total(t *testing.T) {
	testCases := []TestCase{
		{
			total:  0,
			fizzAt: 2,
			buzzAt: 5,
			result: []string{},
		},
	}

	for _, element := range testCases {
		assert.Equal(t, FizzBuzz(element.total, element.fizzAt, element.buzzAt), element.result)
	}
}
