package main

import (
	"aoc21/utils"
	"testing"
)

var exampleInput = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func Test(t *testing.T) {
	utils.StringTest(t, solution, []utils.StringTestData{
		{
			Input:    exampleInput,
			Expected: 150,
		},
	})
}

func Test2(t *testing.T) {
	utils.StringTest(t, solution2, []utils.StringTestData{
		{
			Input:    exampleInput,
			Expected: 900,
		},
	})
}

func TestPart1(t *testing.T) {
	utils.StringTestAoc(t, solution, 2021, 2, 1480518)
}

func TestPart2(t *testing.T) {
	utils.StringTestAoc(t, solution2, 2021, 2, 1282809906)
}
