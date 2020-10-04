package challenge1

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
	panic("implement me")
}
