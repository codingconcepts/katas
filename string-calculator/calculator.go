package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	missingNumberError = errors.New("Missing start or end number")

	digitRegex    = regexp.MustCompile(`^\d$`)
	surroundRegex = regexp.MustCompile(`^\d(?s).*\d$`)
)

type Adder interface {
	Add(numbers string) (result int, err error)
}

type Calculator struct {
	delimiters []rune
}

func NewCalculator(delimiters []rune) (calc *Calculator) {
	return &Calculator{
		delimiters: delimiters,
	}
}

// Add takes a string of comma-separated numbers and
// returns their sum.  It returns zero if nothing is
// provided and an error if a number is missing from
// the start or end.
func (calc *Calculator) Add(numbers string) (result int, err error) {
	numbers = cleanInput(numbers)

	if numbers == "" {
		return 0, nil
	}

	// don't accept incomplete input e.g. "1,"
	if err = validateInput(numbers); err != nil {
		return
	}

	// split and parse numbers to make them easy to sum
	numberInts, err := calc.getNumbers(numbers)
	if err != nil {
		return
	}

	return sum(numberInts), nil
}

// cleanInput removes unnecessary spaces from the
// start and end of the input
func cleanInput(numbers string) (output string) {
	return strings.Trim(numbers, " ")
}

// validateInput performs some sanity checks on the
// raw input to Add.
func validateInput(numbers string) (err error) {
	// only digits at the start and end of the input are allowed
	if !digitRegex.MatchString(numbers) && !surroundRegex.MatchString(numbers) {
		return missingNumberError
	}

	return
}

// getNumbers splits the raw input and parses each
// number found to an integer.
func (calc *Calculator) getNumbers(numbers string) (output []int, err error) {
	output = []int{}

	split := calc.split(numbers)

	for _, numStr := range split {
		if numStr == "" {
			continue
		}

		// trim each
		numStr = strings.Trim(numStr, " ")

		var numInt int
		if numInt, err = strconv.Atoi(numStr); err != nil {
			return
		} else {
			output = append(output, numInt)
		}
	}

	return
}

// split splits a given string input into a collection
// of strings, using the collection of delimiters available
// to the Calculator
func (calc *Calculator) split(numbers string) (split []string) {
	return strings.FieldsFunc(numbers, func(r rune) bool {
		for _, delimiter := range calc.delimiters {
			if r == delimiter {
				return true
			}
		}
		return false
	})
}

// sum iterates through a given slice of numbers
// and adds them to a sum which is returned.
func sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}
