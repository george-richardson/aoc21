package main

import "strconv"

func solution(in *[][]byte) int {
	input := *in
	yMax := len(input)
	xMax := len(input[0])
	var gamma, epsilon []byte
	for x := 0; x < xMax; x++ {
		count := 0
		for y := 0; y < yMax; y++ {
			if input[y][x] == '1' {
				count++
			}
		}
		if count >= yMax / 2 {
			gamma = append(gamma, '1')
			epsilon = append(epsilon, '0')
		} else {
			gamma = append(gamma, '0')
			epsilon = append(epsilon, '1')
		}
	}
	gammai, _ := strconv.ParseInt(string(gamma), 2, 0)
	epsiloni, _ := strconv.ParseInt(string(epsilon), 2, 0)

	return int(gammai * epsiloni)
}

func solution2(in *[][]byte) int {
	oxygenRating, _ := strconv.ParseInt(string(*filter(in, func(ones, zeroes int) byte {
		var r byte = '1'
		if zeroes > ones {
			r = '0'
		}
		return r
	})), 2, 0)
	scrubberRating, _ := strconv.ParseInt(string(*filter(in, func(ones, zeroes int) byte {
		var r byte = '0'
		if zeroes > ones {
			r = '1'
		}
		return r
	})), 2, 0)
	return int(oxygenRating * scrubberRating)
}

type bitCriteria func(int, int)byte

func filter(input *[][]byte, criteria bitCriteria) *[]byte {
	xMax := len((*input)[0])
	for x, yMax := 0, len(*input) ; x < xMax && yMax != 1; x, yMax = x + 1, len(*input) {
		count := 0
		for y := 0; y < yMax; y++ {
			if (*input)[y][x] == '1' {
				count++
			}
		}
		find := criteria(count, yMax - count)
		var newInput [][]byte
		for y := 0; y < yMax; y++ {
			if (*input)[y][x] == find {
				newInput = append(newInput, (*input)[y])
			}
		}
		input = &newInput
	}
	return &(*input)[0]
}
