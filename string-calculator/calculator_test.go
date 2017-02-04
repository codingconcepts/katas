package main

import (
	"fmt"
	"testing"
)

func Test_Q1_Calculator_Add(t *testing.T) {
	testCases := map[string]int{
		"":        0,
		"0":       0,
		"1":       1,
		"10":      10,
		"1,2":     3,
		" 1 , 2 ": 3,
		"1, 2, 3": 6,
	}

	sut := NewCalculator([]rune{','})

	for key, value := range testCases {
		t.Run(fmt.Sprintf("Test_Calculator_Add %v=%v", key, value), func(t *testing.T) {
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

func Test_Q1_Calculator_AddInvalid(t *testing.T) {
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

	sut := NewCalculator([]rune{','})

	for _, value := range testCases {
		t.Run(fmt.Sprintf("Test_Calculator_AddInvalid %q", value), func(t *testing.T) {
			_, err := sut.Add(value)
			if err != missingNumberError {
				t.Fatalf("expected %v but got %v", missingNumberError, err)
			}
		})
	}
}

func Test_Q2_Calculator_AddMoreThan3(t *testing.T) {
	testCases := map[string]int{
		"1, 2, 3, 4":                   10,
		"1, 1, 1, 1, 1, 1, 1, 1, 1, 1": 10,
	}

	sut := NewCalculator([]rune{','})

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

func Test_Q2_Calculator_AddMoreThan3Invalid(t *testing.T) {
	testCases := []string{
		",1, 2, 3, 4",
		"1, 2, 3, 4,",
		",1, 2, 3, 4,",
	}

	sut := NewCalculator([]rune{','})

	for key, value := range testCases {
		t.Run(fmt.Sprintf("Q2_Test_Calculator_AddMoreThan3Invalid %q=%v", key, value), func(t *testing.T) {
			_, err := sut.Add(value)
			if err != missingNumberError {
				t.Fatalf("expected %v but got %v", missingNumberError, err)
			}
		})
	}
}

func Test_Q3_Calculator_Add_WithNewLineDelimiter(t *testing.T) {
	testCases := map[string]int{
		"":          0,
		"0":         0,
		"1":         1,
		"10":        10,
		"1\n2":      3,
		" 1 \n 2 ":  3,
		"1\n 2\n 3": 6,
	}

	sut := NewCalculator([]rune{'\n'})

	for key, value := range testCases {
		t.Run(fmt.Sprintf("Test_Q3_Calculator_Add_WithNewLineDelimiter %v=%v", key, value), func(t *testing.T) {
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

func Test_Q3_Calculator_AddInvalid_WithNewLineDelimiter(t *testing.T) {
	testCases := []string{
		"1\n",
		"\n1",
		"\n1\n",
	}

	sut := NewCalculator([]rune{'\n'})

	for _, value := range testCases {
		t.Run(fmt.Sprintf("Test_Q3_Calculator_AddInvalid_WithNewLineDelimiter %q", value), func(t *testing.T) {
			_, err := sut.Add(value)
			if err != missingNumberError {
				t.Fatalf("expected %v but got %v", missingNumberError, err)
			}
		})
	}
}

func Test_Q3_Calculator_Add_WithCommaAndNewLineDelimiter(t *testing.T) {
	testCases := map[string]int{
		"1\n2,3": 6,
		"1,2\n3": 6,
	}

	sut := NewCalculator([]rune{',', '\n'})

	for key, value := range testCases {
		t.Run(fmt.Sprintf("Test_Q3_Calculator_Add_WithCommaAndNewLineDelimiter %q=%v", key, value), func(t *testing.T) {
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

func Test_Q3_Calculator_AddInvalid_WithCommaAndNewLineDelimiter(t *testing.T) {
	testCases := []string{
		"1,2\n",
		"1\n1,",
	}

	sut := NewCalculator([]rune{',', '\n'})

	for _, value := range testCases {
		t.Run(fmt.Sprintf("Test_Q3_Calculator_AddInvalid_WithCommaAndNewLineDelimiter %q", value), func(t *testing.T) {
			_, err := sut.Add(value)
			if err != missingNumberError {
				t.Fatalf("expected %v but got %v", missingNumberError, err)
			}
		})
	}
}
