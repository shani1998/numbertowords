// Package numberstowords allows numbers to be converted to english words.
package numbertowords

import (
	"errors"
)

var words = [20]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tenwords = [10]string{
	"",
	"",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

// maxNumber is the largest number that this version
// of numberstowords can convert into words.
const maxNumber = 99999

// Convert converts an integer number between 0 and MaxNumber to words.
// If passed a number outside the valid range, it returns an error.
func Convert(number int) (string, error) {
	if number < 0 || number > maxNumber {
		return "", errors.New("number not in valid range")
	}

	result := ""

	thousands := number / 1000
	if thousands > 0 {
		result, _ = Convert(thousands)
		result = result + " thousand "
		number = number % 1000
		if number == 0 {
			return result, nil
		}
	}

	hundreds := number / 100

	if hundreds > 0 {
		result = result + words[hundreds] + " hundred "
		number = number % 100
		if number == 0 {
			return result, nil
		}
	}

	tens := number / 10
	units := number % 10

	if tens < 2 {
		return result + words[number], nil
	}

	if units == 0 {
		return result + tenwords[tens], nil
	}

	return result + tenwords[tens] + " " + words[units], nil

}
