package main

import "strconv"

func solution(in *[]string) int {
	input := *in
	depth := 0
	horizontal := 0
	for _, i := range input {
		x, _ := strconv.Atoi(string(i[len(i)-1]))
		switch i[0] {
		case 'f':
			horizontal += x
		case 'd':
			depth += x
		case 'u':
			depth -= x
		}
	}
	return depth * horizontal
}

func solution2(in *[]string) int {
	input := *in
	depth := 0
	horizontal := 0
	aim := 0
	for _, i := range input {
		x, _ := strconv.Atoi(string(i[len(i)-1]))
		switch i[0] {
		case 'f':
			horizontal += x
			depth += aim * x
		case 'd':
			aim += x
		case 'u':
			aim -= x
		}
	}
	return depth * horizontal
}
