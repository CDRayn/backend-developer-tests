package fizzbuzz

import (
	"strconv"
)

// FizzBuzz performs a FizzBuzz operation over a range of integers
//
// Given a range of integers:
// - Return "Fizz" if the integer is divisible by the `fizzAt` value.
// - Return "Buzz" if the integer is divisible by the `buzzAt` value.
// - Return "FizzBuzz" if the integer is divisible by both the `fizzAt` and
//   `buzzAt` values.
// - Return the original number if is is not divisible by either the `fizzAt` or
//   the `buzzAt` values.
func FizzBuzz(total, fizzAt, buzzAt int64) []string {
	var checkedTotal = total

	// Negative values for total will call make() to panic
	if checkedTotal < 0 {
		return []string{}
	} else if checkedTotal > 1000 {
		// Don't let FizzBuzz() allocate an exorbitant amount of memory using make()
		// truncate it.
		checkedTotal = 1000
	}
	result := make([]string, checkedTotal)

	for i := int64(1); i <= checkedTotal; i++ {
		if fizzAt != 0 && i%fizzAt == 0 {
			result[i-1] += "Fizz"
		}

		if buzzAt != 0 && i%buzzAt == 0 {
			result[i-1] += "Buzz"
		}

		if len(result[i-1]) == 0 {
			result[i-1] = strconv.FormatInt(i, 10)
		}
	}
	return result
}
