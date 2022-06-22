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
	}

	for _, element := range validTestCases {
		assert.Equal(t, FizzBuzz(element.total, element.fizzAt, element.buzzAt), element.result)
	}
}

// TestFizzBuzz_0Fizz checks the behavior of FizzBuzz() for errors (particularly divide by 0 errors) when the fizzAt
// parameter is set to 0.
func TestFizzBuzz_0Fizz(t *testing.T) {
	testCases := []TestCase{
		{
			total:  5,
			fizzAt: 0,
			buzzAt: 3,
			result: []string{"1", "2", buzz, "4", "5"},
		},
		{
			total:  10,
			fizzAt: 0,
			buzzAt: 3,
			result: []string{"1", "2", buzz, "4", "5", buzz, "7", "8", buzz, "10"},
		},
	}

	for _, item := range testCases {
		assert.Equal(t, FizzBuzz(item.total, item.fizzAt, item.buzzAt), item.result)
	}
}

// TestFizzBuzz_TotalLessThan tests the behavior of FizzBuzz() when the total parameter is less than
// either the fizzAt or BuzzAt parameters.
func TestFizzBuzz_TotalLessThan(t *testing.T) {
	testCases := []TestCase{
		{
			total:  2,
			fizzAt: 3,
			buzzAt: 5,
			result: []string{"1", "2"},
		},
		{
			total:  2,
			fizzAt: 3,
			buzzAt: 5,
			result: []string{"1", "2"},
		},
	}

	for _, element := range testCases {
		assert.Equal(t, FizzBuzz(element.total, element.fizzAt, element.buzzAt), element.result)
	}
}

// TestFizzBuzz_0Total tests the behavior of FizzBuzz() when the total parameter is 0. An empty slice
// should be returned.
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
