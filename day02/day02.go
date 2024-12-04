package day02

import (
	"bufio"
	"fmt"
	"leviathan747/aoc24/input"
	"strconv"
	"strings"
)

func Day02() {
	input := input.GetInput("./day02/day02_input.txt")
	reports := ParseInput(input)
	numSafe := NumSafeReports(reports, false)
	fmt.Println(numSafe)
	numSafe = NumSafeReports(reports, true)
	fmt.Println(numSafe)
}

func ParseInput(input string) [][]int {
	reports := [][]int{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		report := []int{}
		for i := 0; i < len(fields); i++ {
			num, _ := strconv.Atoi(fields[i])
			report = append(report, num)
		}
		reports = append(reports, report)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scan error occurred: %v\n", err)
	}
	return reports
}

func ReportIsSafe(report []int) bool {
	sign := 0
	for i := 0; i+1 < len(report); i++ {

		// get the current element, next element, and difference
		curr := report[i]
		next := report[i+1]
		diff := next - curr

		// set the sign for the set based on the first pair
		if sign == 0 {
			if diff < 0 {
				sign = -1
			} else {
				sign = 1
			}
		}

		// check monotonic condition
		if diff < 0 && sign > 0 || diff > 0 && sign < 0 {
			return false
		}

		// check gradual condition
		if diff*sign < 1 || diff*sign > 3 {
			return false
		}

	}
	return true
}

func ReportIsSafeDampened(report []int) bool {
	if ReportIsSafe(report) {
		return true
	} else {
		for i := 0; i < len(report); i++ {
			newReport := []int{}
			newReport = append(newReport, report[:i]...)
			newReport = append(newReport, report[i+1:]...)
			if ReportIsSafe(newReport) {
				return true
			}
		}
		return false
	}
}

func NumSafeReports(reports [][]int, damperEnable bool) int {
	numSafe := 0
	for i := 0; i < len(reports); i++ {
		if damperEnable && ReportIsSafeDampened(reports[i]) || !damperEnable && ReportIsSafe(reports[i]) {
			numSafe++
		}
	}
	return numSafe
}
