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
		if CalibrationIsValid(calibrations[i], []Operation{Multiply, Add}) {
			sum += calibrations[i].test
		}
	}
	fmt.Println(sum)
	sum = 0
	for i := range calibrations {
		if CalibrationIsValid(calibrations[i], []Operation{Multiply, Add, Concat}) {
			sum += calibrations[i].test
		}
	}
	fmt.Println(sum)
}

type Operation func(int, int) int

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

func CalibrationIsValid(c Calibration, operations []Operation) bool {
	for i := 0; i < Pow(len(operations), len(c.operands)-1); i++ {
		k := i
		val := c.operands[0]
		for j := 1; j < len(c.operands); j++ {
			val = operations[k%len(operations)](val, c.operands[j])
			k /= len(operations)
		}
		if val == c.test {
			return true
		}
	}
	return false
}

func Multiply(x, y int) int {
	return x * y
}

func Add(x, y int) int {
	return x + y
}

func Concat(x, y int) int {
	return x*Pow(10, LogBase(y, 10)+1) + y
}

func Pow(x, n int) int {
	if x == 0 || n < 0 {
		return 0
	} else {
		switch n {
		case 0:
			return 1
		case 1:
			return x
		default:
			return Pow(x, n/2) * Pow(x, n-(n/2))
		}
	}
}

func LogBase(x, b int) int {
	val := -1
	for x > 0 {
		x = x / b
		val++
	}
	return val
}
