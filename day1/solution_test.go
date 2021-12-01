package main

import (
	"aoc21/utils"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{199,200,208,210,200,207,240,269,260,263}, 7},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("input %d expect %d", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			actual := solution(&tt.input)
			if actual != tt.expected {
				t.Errorf("actual %d, expected %d", actual, tt.expected)
			}
		})
	}
}

func Test2(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{199,200,208,210,200,207,240,269,260,263}, 5},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("input %d expect %d", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			actual := solution2(&tt.input)
			if actual != tt.expected {
				t.Errorf("actual %d, expected %d", actual, tt.expected)
			}
		})
	}
}

func TestPart1(t *testing.T)  {
	in, err := utils.ReadListOfInts(2021, 1)
	if err != nil {
		t.Error("Error getting input")
	}
	out := solution(in)
	expected := 1387
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}

func TestPart2(t *testing.T)  {
	in, err := utils.ReadListOfInts(2021, 1)
	if err != nil {
		t.Error("Error getting input")
	}
	out := solution2(in)
	expected := 1362
	if out != expected {
		t.Errorf("%d does not match expected %d", out, expected)
	}
	t.Log(out)
}