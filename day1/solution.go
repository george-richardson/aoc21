package main

func solution(in *[]int) int {
	input := *in
	last := input[0]
	count := 0
	for _, i := range input {
		if i > last {
			count++
		}
		last = i
	}
	return count
}

func solution2(in *[]int) int {
	input := *in
	count := 0
	for i := 3; i < len(input); i++ {
		if input[i] > input[i-3] {
			count++
		}
	}
	return count
}