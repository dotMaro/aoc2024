package main

import (
	"fmt"

	"github.com/dotMaro/aoc2024/utils"
)

func main() {
	lines := utils.SplitInput("day04/input.txt")
	lines = lines[:len(lines)-1]
	xmasCount := xmasCount(lines)
	fmt.Println("Part 1.", xmasCount)
	masCount := masCount(lines)
	fmt.Println("Part 2.", masCount)
}

func masCount(lines []string) int {
	count := 0
	for y, line := range lines[:len(lines)-1] {
		if y == 0 {
			continue
		}
		for x, r := range line[:len(line)-1] {
			if x == 0 {
				continue
			}
			if r != 'A' {
				continue
			}

			upLeft := lines[y-1][x-1]
			upRight := lines[y-1][x+1]
			downLeft := lines[y+1][x-1]
			downRight := lines[y+1][x+1]
			if (upLeft == 'M' || upLeft == 'S') &&
				(downRight == 'M' || downRight == 'S') &&
				upLeft != downRight &&
				(upRight == 'M' || upRight == 'S') &&
				(downLeft == 'M' || downLeft == 'S') &&
				upRight != downLeft {
				count++
			}
		}
	}
	return count
}

func xmasCount(lines []string) int {
	count := 0
	for y, line := range lines {
		for x, r := range line {
			if r != 'X' {
				continue
			}
			count += xmasCountOriginatingFromPoint(lines, x, y)
		}
	}
	return count
}

func xmasCountOriginatingFromPoint(lines []string, x int, y int) int {
	if lines[y][x] != 'X' {
		return 0
	}
	count := 0
	for yDelta := -1; yDelta <= 1; yDelta++ {
		for xDelta := -1; xDelta <= 1; xDelta++ {
			if stepFrom(lines, x, y, xDelta, yDelta) {
				count++
			}
		}
	}
	return count
}

func stepFrom(lines []string, x int, y int, xDelta int, yDelta int) bool {
	curY, curX := y, x
	height := len(lines)
	width := len(lines[0])
	for step := range 3 {
		curY += yDelta
		curX += xDelta
		if curY < 0 || curY >= height ||
			curX < 0 || curX >= width {
			return false
		}
		if lines[curY][curX] != "MAS"[step] {
			return false
		}
	}
	return true
}
