package main

import (
	"fmt"
	"testing"
)

func Test_SimpleCalculator_AddValid(t *testing.T) {
	testCases := map[string]int{
		"":        0,
		"0":       0,
		"1":       1,
		"10":      10,
		"1,2":     3,
		" 1 , 2 ": 3,
		"1, 2, 3": 6,
	}

	sut := Calculator{}

	for key, value := range testCases {
		t.Run(fmt.Sprintf("Test_SimpleCalculator_AddValid %q=%v", key, value), func(t *testing.T) {
			actual, err := sut.Add(key)

			if err != nil {
				t.Fatalf("expected no error but got %v", err)
			}
			if actual != value {
				t.Fatalf("expected %v but got %v", value, actual)
			}
		})
	}
}

func Test_SimpleCalculator_AddInvalid(t *testing.T) {
	testCases := []string{
		",",
		",1,",
		"1,",
		",1",
		"1,2,",
		",1,2,",
		",1,2",
		"1,2,3,",
		",1,2,3",
		",1,2,3,",
	}

	sut := Calculator{}

	for _, value := range testCases {
		t.Run(fmt.Sprintf("Test_SimpleCalculator_AddInvalid %q", value), func(t *testing.T) {
			_, err := sut.Add(value)
			if err != missingNumberError {
				t.Fatalf("expected %v but got %v", missingNumberError, err)
			}
		})
	}
}
