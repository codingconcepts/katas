package calc

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	missingNumberError = errors.New("Missing start or end number")

	digitRegex    = regexp.MustCompile(`^-?\d$`)
	surroundRegex = regexp.MustCompile(`^-?\d[\s\S]*-?\d$`)

	// TODO: this regex doesn't allow new line characters to
	// be used as delimiters at the moment
	customDelimiterRegex = regexp.MustCompile(`//([\s\S]*)\n`)
)

type Adder interface {
	Add(numbers string) (result int, err error)
}

type Calculator struct {
	delimiters []rune
}

// NegativeNumberError is encountered when a negative
// number is passed to the Calculator's Add method.
type NegativeNumberError struct {
	Numbers []int
}

// NewCalculator spins up a pointer to a new Calculator.
func NewCalculator(delimiters ...rune) (calc *Calculator) {
	return &Calculator{
		delimiters: delimiters,
	}
}

// NewNegativeNumberError creates the error with a
// collection containing the negative numbers that
// caused the error.
func NewNegativeNumberError(numbers []int) error {
	neg := NegativeNumberError{}

	neg.Numbers = make([]int, len(numbers))
	copy(neg.Numbers, numbers)

	return neg
}

func (e NegativeNumberError) Error() string {
	return fmt.Sprintf("negatives not allowed: %v", e.Numbers)
}

// Add takes a string of comma-separated numbers and
// returns their sum.  It returns zero if nothing is
// provided and an error if a number is missing from
// the start or end.
func (calc *Calculator) Add(numbers string) (result int, err error) {
	numbers = cleanInput(numbers)

	// parse a customer delimiter if one has been provided,
	// then strip it from the input.  doing this here and
	// not in the Calculator setup because the custom
	// delimiter is provided via the input to Add, not to
	// the calculator itself
	if calc.parseCustomDelimeter(numbers) {
		numbers = calc.stripCustomerDelimiter(numbers)
	}

	if numbers == "" {
		return 0, nil
	}

	// validate the input string provided by the user,
	// don't accept incomplete input e.g. "1,"
	if err = validateInput(numbers); err != nil {
		return
	}

	// split and parse numbers to make them easy to sum
	numberInts, err := calc.getNumbers(numbers)
	if err != nil {
		return
	}

	// validate the individual numbers provided by the user,
	// don't accept negative input
	if err = calc.validateNumbers(numberInts); err != nil {
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

// validateNumbers validates each of the user's
// provided numbers
func (calc *Calculator) validateNumbers(numbers []int) (err error) {
	neg := []int{}
	for _, number := range numbers {
		if number < 0 {
			neg = append(neg, number)
		}
	}

	if len(neg) > 0 {
		return NewNegativeNumberError(neg)
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

		// trim each number
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

// parseCustomDelimeter allows the user to provide
// a customer split delimiter between numbers.
func (calc *Calculator) parseCustomDelimeter(numbers string) (ok bool) {
	if !customDelimiterRegex.MatchString(numbers) {
		return false
	}

	results := customDelimiterRegex.FindStringSubmatch(numbers)
	if len(results) == 0 {
		return
	}

	// replace the calculator's delimiters with the
	// user's new one
	calc.delimiters = []rune{
		rune(results[1][0]),
	}

	return true
}

// stripCustomerDelimiter removes the user's custom
// delimeter from the input, allowing the result of
// it to be parsed.
func (*Calculator) stripCustomerDelimiter(numbers string) (output string) {
	return customDelimiterRegex.ReplaceAllString(numbers, "")
}
