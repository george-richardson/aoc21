package utils

import (
	"bufio"
	"fmt"
	"testing"
)

// Strings

type StringSolution func(*[]string) int
type StringTestData struct {
	Input    []string
	Expected int
}

func StringTests(t *testing.T, solution StringSolution, testData []StringTestData) {
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

func StringTest(t *testing.T, solution StringSolution, testData StringTestData) {
	testName := fmt.Sprintf("input %v expect %v", testData.Input, testData.Expected)
	t.Run(testName, func(t *testing.T) {
		actual := solution(&testData.Input)
		if actual != testData.Expected {
			t.Errorf("actual %v, expected %v", actual, testData.Expected)
		}
	})
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

// Grid of bytes

type GridOfBytesSolution func(*[][]byte) int
type GridOfBytesTestData struct {
	Input    [][]byte
	Expected int
}

func GridOfBytesTest(t *testing.T, solution GridOfBytesSolution, testData GridOfBytesTestData) {
	testName := fmt.Sprintf("input %v expect %v", testData.Input, testData.Expected)
	t.Run(testName, func(t *testing.T) {
		actual := solution(&testData.Input)
		if actual != testData.Expected {
			t.Errorf("actual %v, expected %v", actual, testData.Expected)
		}
	})
}

func GridOfBytesTests(t *testing.T, solution GridOfBytesSolution, testData []GridOfBytesTestData) {
	for _, tt := range testData {
		GridOfBytesTest(t, solution, tt)
	}
}

func GridOfBytesTestAoc(t *testing.T, solution GridOfBytesSolution, year int, day int, expected int) {
	in, err := ReadGridOfBytes(year, day)
	if err != nil {
		t.Error("Error getting input")
	}
	out := solution(in)
	if out != expected {
		t.Errorf("%v does not match expected %v", out, expected)
	}
	t.Log(out)
}

// Strings

type ScannerSolution func(scanner *bufio.Scanner) int
type ScannerTestData struct {
	Input    *bufio.Scanner
	Expected int
}

func ScannerTests(t *testing.T, solution ScannerSolution, testData []ScannerTestData) {
	for _, tt := range testData {
		testName := fmt.Sprintf("input %v expect %v", tt.Input, tt.Expected)
		t.Run(testName, func(t *testing.T) {
			actual := solution(tt.Input)
			if actual != tt.Expected {
				t.Errorf("actual %v, expected %v", actual, tt.Expected)
			}
		})
	}
}

func ScannerTest(t *testing.T, solution ScannerSolution, testData ScannerTestData) {
	testName := fmt.Sprintf("input %v expect %v", testData.Input, testData.Expected)
	t.Run(testName, func(t *testing.T) {
		actual := solution(testData.Input)
		if actual != testData.Expected {
			t.Errorf("actual %v, expected %v", actual, testData.Expected)
		}
	})
}

func ScannerTestAoc(t *testing.T, solution ScannerSolution, year int, day int, expected int) {
	in, err := GetInputScanner(year, day)
	if err != nil {
		t.Error("Error getting input")
	}
	out := solution(in)
	if out != expected {
		t.Errorf("%v does not match expected %v", out, expected)
	}
	t.Log(out)
}
