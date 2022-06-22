package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzzValidEntries(t *testing.T) {

	const fizz = "Fizz"
	const buzz = "Buzz"
	const fizzbuzz = "FizzBuzz"

	result := []string{"1", fizz, "3", fizz, buzz, fizz, "7", fizz, "9", fizzbuzz}

	assert.Equal(t, FizzBuzz(10, 2, 5), result)
}
