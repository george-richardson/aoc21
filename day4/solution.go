package main

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type bingoBoard struct {
	grid   [5][5]int
	marks  [5][5]bool
	xMarks [5]int
	yMarks [5]int
	won    bool
}

func (b *bingoBoard) mark(num int) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if b.grid[y][x] == num {
				b.marks[y][x] = true
				b.xMarks[x] = b.xMarks[x] + 1
				b.yMarks[y] = b.yMarks[y] + 1
			}
		}
	}
}

func (b *bingoBoard) isWinner() bool {
	for i := 0; i < 5; i++ {
		if b.xMarks[i] == 5 || b.yMarks[i] == 5 {
			return true
		}
	}
	return false
}

func (b *bingoBoard) justWon() bool {
	if b.won {
		return false
	}
	if b.isWinner() {
		b.won = true
		return true
	}
	return false
}

func (b *bingoBoard) calulateScore() (score int) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !b.marks[y][x] {
				score += b.grid[y][x]
			}
		}
	}
	return score
}

func readBoard(scanner *bufio.Scanner) *bingoBoard {
	board := bingoBoard{}
	for y := 0; y < 5; y++ {
		scanner.Scan()
		line := strings.Trim(scanner.Text(), " ")
		for x, i := range regexp.MustCompile("\\s+").Split(line, -1) {
			num, _ := strconv.Atoi(i)
			board.grid[4-y][x] = num
		}
	}
	return &board
}

func readBoards(scanner *bufio.Scanner) (boards []*bingoBoard, draws []int) {
	scanner.Scan()
	line := scanner.Text()
	for _, i := range strings.Split(line, ",") {
		num, _ := strconv.Atoi(i)
		draws = append(draws, num)
	}

	for scanner.Scan() {
		boards = append(boards, readBoard(scanner))
	}
	return
}

func solution(in *bufio.Scanner) int {
	boards, draws := readBoards(in)
	for _, draw := range draws {
		for _, board := range boards {
			board.mark(draw)
		}
		for _, board := range boards {
			if board.isWinner() {
				return board.calulateScore() * draw
			}
		}
	}
	return -1
}

func solution2(in *bufio.Scanner) int {
	boards, draws := readBoards(in)
	winningBoards := 0
	totalBoards := len(boards)
	for _, draw := range draws {
		for _, board := range boards {
			board.mark(draw)
		}
		for _, board := range boards {
			if board.justWon() {
				winningBoards++
				if winningBoards == totalBoards {
					return board.calulateScore() * draw
				}
			}
		}
	}
	return -1
}
