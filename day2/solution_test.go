package main

import (
	"aoc21/utils"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{[]string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}, 150},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("input %v expect %v", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			actual := solution(&tt.input)
			if actual != tt.expected {
				t.Errorf("actual %v, expected %v", actual, tt.expected)
			}
		})
	}
}

func Test2(t *testing.T) {
	var tests = []struct {
		input    []string
		expected int
	}{
		{[]string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}, 900},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("input %v expect %v", tt.input, tt.expected)
		t.Run(testName, func(t *testing.T) {
			actual := solution2(&tt.input)
			if actual != tt.expected {
				t.Errorf("actual %v, expected %v", actual, tt.expected)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	in, err := utils.ReadListOfStrings(2021, 2)
	if err != nil {
		t.Error("Error getting input")
	}
	out := solution(in)
	expected := 1480518
	if out != expected {
		t.Errorf("%v does not match expected %v", out, expected)
	}
	t.Log(out)
}

func TestPart2(t *testing.T) {
	in, err := utils.ReadListOfStrings(2021, 2)
	if err != nil {
		t.Error("Error getting input")
	}
	out := solution2(in)
	expected := 1282809906
	if out != expected {
		t.Errorf("%v does not match expected %v", out, expected)
	}
	t.Log(out)
}
