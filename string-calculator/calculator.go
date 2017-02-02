package main

import (
	"errors"
	"strconv"
	"strings"
)

var (
	missingNumberError = errors.New("Missing start or end number")
)

type Adder interface {
	Add(numbers string) (result int, err error)
}

type Calculator struct {
}

// Add takes a string of comma-separated numbers and
// returns their sum.  It returns zero if nothing is
// provided and an error if a number is missing from
// the start or end.
func (calc *Calculator) Add(numbers string) (result int, err error) {
	if numbers == "" {
		return 0, nil
	}

	// don't accept incomplete input e.g. "1,"
	if err = validateInput(numbers); err != nil {
		return
	}

	// split and parse numbers to make them easy to sum
	numberInts, err := getNumbers(numbers)
	if err != nil {
		return
	}

	return sum(numberInts), nil
}

// validateInput performs some sanity checks on the
// raw input to Add.
func validateInput(numbers string) (err error) {
	trimmed := strings.Trim(numbers, " ")

	if strings.HasPrefix(trimmed, ",") || strings.HasSuffix(trimmed, ",") {
		return missingNumberError
	}

	return
}

// getNumbers splits the raw input and parses each
// number found to an integer.
func getNumbers(numbers string) (output []int, err error) {
	output = []int{}

	clean := strings.Replace(numbers, " ", "", -1)
	split := strings.Split(clean, ",")

	for _, numStr := range split {
		if numStr == "" {
			continue
		}

		var numInt int
		if numInt, err = strconv.Atoi(numStr); err != nil {
			return
		} else {
			output = append(output, numInt)
		}
	}

	return
}

// sum iterates through a given slice of numbers
// and adds them to a sum which is returned.
func sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}
