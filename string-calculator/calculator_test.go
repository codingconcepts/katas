package main

import (
	"fmt"
	"testing"
)

func Q1_Test_Calculator_Add(t *testing.T) {
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
		t.Run(fmt.Sprintf("Test_Calculator_Add %q=%v", key, value), func(t *testing.T) {
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

func Q1_Test_Calculator_AddInvalid(t *testing.T) {
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
		t.Run(fmt.Sprintf("Test_Calculator_AddInvalid %q", value), func(t *testing.T) {
			_, err := sut.Add(value)
			if err != missingNumberError {
				t.Fatalf("expected %v but got %v", missingNumberError, err)
			}
		})
	}
}

func Q2_Test_Calculator_AddMoreThan3(t *testing.T) {
	testCases := map[string]int{
		"1, 2, 3, 4":                   10,
		"1, 1, 1, 1, 1, 1, 1, 1, 1, 1": 10,
	}

	sut := Calculator{}

	for key, value := range testCases {
		t.Run(fmt.Sprintf("Q2_Test_Calculator_AddMoreThan3 %q=%v", key, value), func(t *testing.T) {
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

func Q2_Test_Calculator_AddMoreThan3Invalid(t *testing.T) {
	testCases := []string{
		",1, 2, 3, 4",
		"1, 2, 3, 4,",
		",1, 2, 3, 4,",
	}

	sut := Calculator{}

	for key, value := range testCases {
		t.Run(fmt.Sprintf("Q2_Test_Calculator_AddMoreThan3Invalid %q=%v", key, value), func(t *testing.T) {
			_, err := sut.Add(value)
			if err != missingNumberError {
				t.Fatalf("expected %v but got %v", missingNumberError, err)
			}
		})
	}
}
