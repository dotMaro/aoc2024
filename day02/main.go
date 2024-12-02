package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/dotMaro/aoc2024/utils"
)

func main() {
	input := utils.SplitInput("day02/input.txt")
	reports := make([]report, len(input)-1)
	for i, line := range input[:len(input)-1] {
		reports[i] = parseReport(line)
	}
	safeCount := 0
	for _, report := range reports {
		if report.isSafe() {
			safeCount++
		}
	}
	fmt.Println("Part 1.", safeCount)
	safeByOmittingCount := 0
	for _, report := range reports {
		if report.isSafeByOmittingAnyOneLevel() {
			safeByOmittingCount++
		}
	}
	fmt.Println("Part 2.", safeByOmittingCount)
}

type report []int

func parseReport(text string) report {
	splitText := strings.Split(text, " ")
	report := make([]int, len(splitText))
	for i, number := range splitText {
		var err error
		report[i], err = strconv.Atoi(number)
		if err != nil {
			fmt.Println(text)
			panic(err)
		}
	}
	return report
}

func (r report) isSafe() bool {
	isIncreasing := r[1] > r[0]
	lastLevel := r[0]
	for _, level := range r[1:] {
		increasedComparedToLast := level > lastLevel
		if increasedComparedToLast != isIncreasing {
			return false
		}
		diff := level - lastLevel
		if diff < 0 {
			diff *= -1
		}
		if diff < 1 || diff > 3 {
			return false
		}
		lastLevel = level
	}
	return true
}

func (r report) isSafeByOmittingAnyOneLevel() bool {
	if r.isSafe() {
		return true
	}
	for i := range len(r) {
		omittedReport := slices.Clone(r)
		omittedReport = append(omittedReport[:i], omittedReport[i+1:]...)
		if omittedReport.isSafe() {
			return true
		}
	}
	return false
}
