package day06

import (
	"bufio"
	"fmt"
	"leviathan747/aoc24/input"
	"strings"
)

func Day06() {
	input := input.GetInput("./day06/day06_input.txt")
	theMap, x, y, xInc, yInc := ParseInput(input)
	steps := CountSteps(theMap, x, y, xInc, yInc)
	fmt.Println(steps)
	obstructions := CountObstructionLocations(theMap, x, y, xInc, yInc)
	fmt.Println(obstructions)
}

type RoomMap [][]bool

func (r RoomMap) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[i]); j++ {
			if r[i][j] {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func ParseInput(input string) (RoomMap, int, int, int, int) {

	theMap := RoomMap{}
	x, y, xInc, yInc := 0, 0, 0, 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		row := make([]bool, len(line))
		theMap = append(theMap, row)
		for j, c := range scanner.Text() {
			switch c {
			case '#':
				row[j] = true
			case '^':
				x, y = j, i
				xInc, yInc = 0, -1
			case 'v':
				x, y = j, i
				xInc, yInc = 0, 1
			case '<':
				x, y = j, i
				xInc, yInc = -1, 0
			case '>':
				x, y = j, i
				xInc, yInc = 1, 0
			}

		}
	}

	return theMap, x, y, xInc, yInc

}

func RotateRight(xInc, yInc int) (int, int) {
	if xInc == 0 && yInc == -1 {
		return 1, 0
	} else if xInc == 1 && yInc == 0 {
		return 0, 1
	} else if xInc == 0 && yInc == 1 {
		return -1, 0
	} else if xInc == -1 && yInc == 0 {
		return 0, -1
	} else {
		return xInc, yInc
	}
}

func MakeMove(theMap RoomMap, x int, y int, xInc int, yInc int) (int, int, int, int) {
	if !theMap[y+yInc][x+xInc] { // take a step
		return x + xInc, y + yInc, xInc, yInc
	} else { // turn to the right
		xInc, yInc = RotateRight(xInc, yInc)
		return x, y, xInc, yInc
	}
}

func CountSteps(theMap RoomMap, x int, y int, xInc int, yInc int) int {

	steps := 0
	path := map[[2]int]bool{}

	// break out when we reach the edge
	for x > 0 && x < len(theMap[0])-1 && y > 0 && y < len(theMap)-1 {
		x, y, xInc, yInc = MakeMove(theMap, x, y, xInc, yInc)
		if _, present := path[[2]int{x, y}]; !present {
			steps++
			path[[2]int{x, y}] = true
		}
	}

	return steps

}

func IsLoop(theMap RoomMap, x int, y int, xInc int, yInc int) bool {

	steps := 0
	path := map[[2]int]bool{}
	pathWithDirection := map[[4]int]bool{}

	// break out when we reach the edge
	for steps < len(theMap)*len(theMap[0]) && x > 0 && x < len(theMap[0])-1 && y > 0 && y < len(theMap)-1 {
		x, y, xInc, yInc = MakeMove(theMap, x, y, xInc, yInc)
		if _, present := path[[2]int{x, y}]; !present {
			steps++
			path[[2]int{x, y}] = true
		}
		if _, present := pathWithDirection[[4]int{x, y, xInc, yInc}]; present {
			return true
		}
		pathWithDirection[[4]int{x, y, xInc, yInc}] = true
	}

	return false

}

func CountObstructionLocations(theMap RoomMap, x int, y int, xInc int, yInc int) int {

	obstructions := 0

	for i := 0; i < len(theMap); i++ {
		for j := 0; j < len(theMap[i]); j++ {
			if (j != x || i != y) && !theMap[i][j] {
				newMap := make(RoomMap, len(theMap))
				for k := 0; k < len(theMap); k++ {
					newMap[k] = make([]bool, len(theMap[k]))
					copy(newMap[k], theMap[k])
				}
				newMap[i][j] = true
				if IsLoop(newMap, x, y, xInc, yInc) {
					obstructions++
				}
			}
		}
	}

	return obstructions

}
