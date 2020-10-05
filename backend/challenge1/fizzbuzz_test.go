package challenge1_test

import (
	"testing"

	"github.com/Clarilab/codechallenge/backend/challenge1"
)

func Test_FizzBuzz_Fizz(t *testing.T) {
	// Arrange
	fizzBuzzer := challenge1.NewFizzBuzzer()

	// Act
	result, err := fizzBuzzer.FizzBuzz(3)
	if err != nil {
		t.Error(err)
	}

	// Assert
	if result != "Fizz" {
		t.Error("Expected result: Fizz", "actual result:", result)
	}
}

func Test_FizzBuzz_Buzz(t *testing.T) {
	// Arrange
	fizzBuzzer := challenge1.NewFizzBuzzer()

	// Act
	result, err := fizzBuzzer.FizzBuzz(5)
	if err != nil {
		t.Error(err)
	}

	// Assert
	if result != "Buzz" {
		t.Error("Expected result: Buzz", "actual result:", result)
	}
}

func Test_FizzBuzz_FizzBuzz(t *testing.T) {
	// Arrange
	fizzBuzzer := challenge1.NewFizzBuzzer()

	// Act
	result, err := fizzBuzzer.FizzBuzz(15)
	if err != nil {
		t.Error(err)
	}

	// Assert
	if result != "FizzBuzz" {
		t.Error("Expected result: FizzBuzz", "actual result:", result)
	}
}

func Test_FizzBuzz_EmptyString(t *testing.T) {
	// Arrange
	fizzBuzzer := challenge1.NewFizzBuzzer()

	// Act
	result, err := fizzBuzzer.FizzBuzz(43)
	if err != nil {
		t.Error(err)
	}

	// Assert
	if result != "" {
		t.Error("Expected result: ", "actual result:", result)
	}
}

func Test_FizzBuzz_Error(t *testing.T) {
	// Arrange
	fizzBuzzer := challenge1.NewFizzBuzzer()

	// Act
	result, err := fizzBuzzer.FizzBuzz(-1337)
	if err == nil {
		t.Error("Expected error to be not nil")
	}

	// Assert
	if result != "" {
		t.Error("Expected result: ", "actual result:", result)
	}
}
