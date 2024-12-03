package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dotMaro/aoc2024/utils"
)

var mulInstructions *regexp.Regexp

func init() {
	mulInstructions = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
}

func main() {
	input := utils.InputString("day03/input.txt")
	matches := mulInstructions.FindAllStringIndex(input, -1)
	sum := 0
	sumWithConditional := 0
	for _, m := range matches {
		matchBegin, matchEnd := m[0], m[1]
		instruction := parseInstruction(input[matchBegin:matchEnd])
		mulResult := instruction.multiply()
		sum += mulResult
		if mulInstructionEnabledAt(input, matchBegin-len("do()")) {
			sumWithConditional += mulResult
		}
	}
	fmt.Println("Part 1.", sum)
	fmt.Println("Part 2.", sumWithConditional)
}

func mulInstructionEnabledAt(text string, index int) bool {
	for {
		if index < 0 || strings.HasPrefix(text[index:], "do()") {
			return true
		}
		if strings.HasPrefix(text[index:], "don't()") {
			return false
		}
		index--
	}
}

type instruction struct {
	left, right int
}

func parseInstruction(text string) instruction {
	text = text[len("mul(") : len(text)-len(")")]
	split := strings.Split(text, ",")
	left, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	return instruction{
		left:  left,
		right: right,
	}
}

func (i instruction) multiply() int {
	return i.left * i.right
}
