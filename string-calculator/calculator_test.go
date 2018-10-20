package calc

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Q1_Calculator_AddEmpty(t *testing.T) {
	sut := NewCalculator(',')

	actual, err := sut.Add("")

	if err != nil {
		t.Fatalf("expected no error but got %v", err)
	}
	if actual != 0 {
		t.Fatalf("expected %v but got %v", 0, actual)
	}
}

func Test_Q1_Calculator_Add(t *testing.T) {
	testCases := map[string]int{
		"0":       0,
		"1":       1,
		"10":      10,
		"1,2":     3,
		" 1 , 2 ": 3,
		"1, 2, 3": 6,
	}

	sut := NewCalculator(',')

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

	sut := NewCalculator(',')

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

	sut := NewCalculator(',')

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

	sut := NewCalculator(',')

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

	sut := NewCalculator('\n')

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

	sut := NewCalculator('\n')

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

	sut := NewCalculator(',', '\n')

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

	sut := NewCalculator(',', '\n')

	for _, value := range testCases {
		t.Run(fmt.Sprintf("Test_Q3_Calculator_AddInvalid_WithCommaAndNewLineDelimiter %q", value), func(t *testing.T) {
			_, err := sut.Add(value)
			if err != missingNumberError {
				t.Fatalf("expected %v but got %v", missingNumberError, err)
			}
		})
	}
}

func Test_Q4_Calculator_Add_WithCustomDelimiter(t *testing.T) {
	testCases := map[string]int{
		"//;\n1":     1,
		"//;\n1;2":   3,
		"//;\n1;2;3": 6,
	}

	sut := NewCalculator()

	for key, value := range testCases {
		t.Run(fmt.Sprintf("Test_Q4_Calculator_Add_WithCustomDelimiter %q=%v", key, value), func(t *testing.T) {
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

func Test_Q4_Calculator_AddInvalid_WithCustomDelimiter(t *testing.T) {
	testCases := []string{
		"//;\n1;",
		"//;\n1;2;",
		"//;\n1;2;3;",
	}

	sut := NewCalculator()

	for _, value := range testCases {
		t.Run(fmt.Sprintf("Test_Q4_Calculator_AddInvalid_WithCustomDelimiter %q", value), func(t *testing.T) {
			_, err := sut.Add(value)
			if err != missingNumberError {
				t.Fatalf("expected %v but got %v", missingNumberError, err)
			}
		})
	}
}

func Test_Q5_AddNegative(t *testing.T) {
	testCases := []struct {
		input          string
		negativeInputs []int
	}{
		{input: "-1", negativeInputs: []int{-1}},
		{input: "-1, 2", negativeInputs: []int{-1}},
		{input: "-1, 2, -3", negativeInputs: []int{-1, -3}},
		{input: "-1\n-2\n-3", negativeInputs: []int{-1, -2, -3}},
		{input: "//;\n-1;-2;-4", negativeInputs: []int{-1, -2, -4}},
		{input: "//;\n-1;-2;-3; 4; 5; -6", negativeInputs: []int{-1, -2, -3, -6}},
	}

	sut := NewCalculator(',', '\n')

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test_Q5_AddNegative %q", testCase.input), func(t *testing.T) {
			_, err := sut.Add(testCase.input)
			if err == nil {
				t.Fatal("Expected an error but didn't get one")
			}

			actual, ok := err.(NegativeNumberError)
			if !ok {
				t.Fatalf("Expected NegativeNumberError but got %q", actual)
			}

			if !reflect.DeepEqual(testCase.negativeInputs, actual.Numbers) {
				t.Fatalf("Expected %v but got %q", testCase.negativeInputs, actual.Numbers)
			}
		})
	}
}

func Test_Q5_Calculator_AddNegativeInvalid(t *testing.T) {
	testCases := []string{
		"-1,",
		",-1",
		",-1,",
	}

	sut := NewCalculator()

	for _, value := range testCases {
		t.Run(fmt.Sprintf("Test_Q5_Calculator_AddNegativeInvalid %q", value), func(t *testing.T) {
			_, err := sut.Add(value)
			if err != missingNumberError {
				t.Fatalf("expected %v but got %v", missingNumberError, err)
			}
		})
	}
}
