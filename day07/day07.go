package day07

import (
	"bufio"
	"fmt"
	"leviathan747/aoc24/input"
	"regexp"
	"strconv"
	"strings"
)

func Day07() {
	input := input.GetInput("./day07/day07_input.txt")
	calibrations := ParseInput(input)
	sum := 0
	for i := range calibrations {
		if CalibrationIsValid(calibrations[i]) {
			sum += calibrations[i].test
		}
	}
	fmt.Println(sum)
}

type Calibration struct {
	test     int
	operands []int
}

func ParseInput(input string) []Calibration {

	calibrations := []Calibration{}

	re := regexp.MustCompile(`([1-9][0-9]*):\s+(([1-9][0-9]*\s*)+)`)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		m := re.FindStringSubmatch(line)
		if m != nil {
			if n, err := strconv.Atoi(m[1]); err == nil {
				calibrations = append(calibrations, Calibration{n, StringsToInts(strings.Fields(m[2]))})
			}
		}
	}

	return calibrations
}

func StringsToInts(input []string) []int {
	result := make([]int, len(input))
	for i := range input {
		if n, err := strconv.Atoi(input[i]); err == nil {
			result[i] = n
		}
	}
	return result
}

func CalibrationIsValid(c Calibration) bool {
	for i := 0; i < 0b1<<(len(c.operands)-1); i++ {
		val := c.operands[0]
		for j := 1; j < len(c.operands); j++ {
			if (i>>(j-1))&0b1 == 1 {
				val *= c.operands[j]
			} else {
				val += c.operands[j]
			}
		}
		if val == c.test {
			return true
		}
	}
	return false
}
