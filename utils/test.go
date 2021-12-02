package utils

import (
	"fmt"
	"testing"
)

type StringSolution func(*[]string) int
type StringTestData struct {
	Input    []string
	Expected int
}

func StringTest(t *testing.T, solution StringSolution, testData []StringTestData) {
	for _, tt := range testData {
		testName := fmt.Sprintf("input %v expect %v", tt.Input, tt.Expected)
		t.Run(testName, func(t *testing.T) {
			actual := solution(&tt.Input)
			if actual != tt.Expected {
				t.Errorf("actual %v, expected %v", actual, tt.Expected)
			}
		})
	}
}

func StringTestAoc(t *testing.T, solution StringSolution, year int, day int, expected int) {
	in, err := ReadListOfStrings(year, day)
	if err != nil {
		t.Error("Error getting input")
	}
	out := solution(in)
	if out != expected {
		t.Errorf("%v does not match expected %v", out, expected)
	}
	t.Log(out)
}
