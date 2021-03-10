package numbertowords

import (
	"errors"
)

// MaxNumber is the largest number that this version
// of numbertowords can convert to words.
const (
	MaxNumber = 99999
)

func Convert(num int) (string, error) {
	if num < 0 || num > MaxNumber {
		return "", errors.New("number not in a valid range")
	}
	unitWords := [20]string{
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

	tenWords := [10]string{
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

	result := ""

	thousands := num / 1000

	if thousands > 0 {
		result = unitWords[thousands] + " thousands "
		num = num % 1000
		if num == 0 {
			return result, nil
		}

	}

	hundreds := num / 100

	if hundreds > 0 {
		result = result + unitWords[hundreds] + " hundreds "
		num = num % 100
		if num == 0 {
			return result, nil
		}
	}

	tens := num / 10
	units := num % 10

	if tens < 2 {
		return result + unitWords[num], nil
	}
	if units == 0 {
		return result + tenWords[tens], nil
	}
	return result + tenWords[tens] + " " + unitWords[units], nil

}
