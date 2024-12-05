package main

import (
	"testing"

	"github.com/dotMaro/aoc2024/utils"
)

func Test_xmasCount(t *testing.T) {
	const input = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	lines := utils.SplitLine(input)
	got := xmasCount(lines)
	if got != 18 {
		t.Errorf("xmasCount() = %d; want 18", got)
	}
}
