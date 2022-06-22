package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const fizz = "Fizz"
const buzz = "Buzz"
const fizzbuzz = "FizzBuzz"

func TestFizzBuzzValidEntries(t *testing.T) {
	result := []string{"1", fizz, "3", fizz, buzz, fizz, "7", fizz, "9", fizzbuzz}

	assert.Equal(t, FizzBuzz(10, 2, 5), result)
}
