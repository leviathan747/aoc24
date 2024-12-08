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

func CountSteps(theMap RoomMap, x int, y int, xInc int, yInc int) int {

	steps := 0
	path := map[int]bool{}

	// break out when we reach the edge
	for x > 0 && x < len(theMap[0])-1 && y > 0 && y < len(theMap)-1 {
		if !theMap[y+yInc][x+xInc] { // take a step
			x, y = x+xInc, y+yInc
			if _, present := path[y*len(theMap)+x]; !present {
				steps++
				path[y*len(theMap)+x] = true
			}
		} else { // turn to the right
			if xInc == 0 && yInc == -1 {
				xInc, yInc = 1, 0
			} else if xInc == 1 && yInc == 0 {
				xInc, yInc = 0, 1
			} else if xInc == 0 && yInc == 1 {
				xInc, yInc = -1, 0
			} else if xInc == -1 && yInc == 0 {
				xInc, yInc = 0, -1
			}
		}
	}

	return steps

}
