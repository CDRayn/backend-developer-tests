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

//// TestFizzBuzz_0Fizz tests the behavior of FizzBuzz() when the fizzAt param is 0. This should raise an error
//// to avoid a divide by zero bug
//func TestFizzBuzz_0Fizz(t *testing.T) {
//	testCases := []TestCase{
//		// Test for divide by zero errors
//		{
//			total:  5,
//			fizzAt: 0,
//			buzzAt: 3,
//			result: []string{"1", "2", buzz, "4", "5"},
//		},
//	}
//
//	for _, element := range testCases {
//		_, err := FizzBuzz(element.total, element.fizzAt, element.total)
//		assert.Error(t, err, "Error expected but no encounted")
//
//	}
//}

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
