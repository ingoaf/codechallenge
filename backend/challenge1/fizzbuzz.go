package challenge1

import "errors"

const Fizz = "Fizz"
const Buzz = "Buzz"

// FizzBuzzer provides FizzBuzz functionality
type FizzBuzzer interface {
	FizzBuzz(number int) (string, error)
}

type fizzBuzz struct{}

// NewFizzBuzzer creates a new instance of FizzBuzzer
func NewFizzBuzzer() FizzBuzzer {
	return &fizzBuzz{}
}

func (f *fizzBuzz) FizzBuzz(number int) (string, error) {
	if number < 0 {
		return "", errors.New("Negative numbers are not allowed")
	}

	var answer string
	if number%3 == 0 {
		answer += Fizz
	}
	if number%5 == 0 {
		answer += Buzz
	}
	return answer, nil
}
