package main

import (
	"aoc21/utils"
	"testing"
)

var exampleInput = [][]byte{
	{ '0', '0', '1', '0', '0' },
	{ '1', '1', '1', '1', '0' },
	{ '1', '0', '1', '1', '0' },
	{ '1', '0', '1', '1', '1' },
	{ '1', '0', '1', '0', '1' },
	{ '0', '1', '1', '1', '1' },
	{ '0', '0', '1', '1', '1' },
	{ '1', '1', '1', '0', '0' },
	{ '1', '0', '0', '0', '0' },
	{ '1', '1', '0', '0', '1' },
	{ '0', '0', '0', '1', '0' },
	{ '0', '1', '0', '1', '0' },
}

func Test(t *testing.T) {
	utils.GridOfBytesTest(t, solution, utils.GridOfBytesTestData{ Input: exampleInput, Expected: 198 })
}

func Test2(t *testing.T) {
	utils.GridOfBytesTest(t, solution2, utils.GridOfBytesTestData { Input: exampleInput, Expected: 230 })
}

func TestPart1(t *testing.T) {
	utils.GridOfBytesTestAoc(t, solution, 2021, 3, 1458194)
}

func TestPart2(t *testing.T) {
	utils.GridOfBytesTestAoc(t, solution2, 2021, 3, 2829354)
}
