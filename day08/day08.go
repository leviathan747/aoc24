package day08

import (
	"bufio"
	"fmt"
	"leviathan747/aoc24/input"
	"regexp"
	"strings"
)

func Day08() {
	input := input.GetInput("./day08/day08_input.txt")
	amap, width, height := ParseInput(input)
	antiNodeMap := BuildAntiNodeMap(amap, width, height, false)
	fmt.Println(len(antiNodeMap))
}

type AntennaMap map[rune][][2]int
type AntiNodeMap map[[2]int]int

func ParseInput(input string) (AntennaMap, int, int) {
	amap := AntennaMap{}

	re := regexp.MustCompile(`[0-9a-zA-Z]`)
	scanner := bufio.NewScanner(strings.NewReader(input))
	i := 0
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		matches := re.FindAllStringIndex(line, -1)
		for n := range matches {
			c := rune(line[matches[n][0]])
			if _, present := amap[c]; present {
				amap[c] = append(amap[c], [2]int{i, matches[n][0]})
			} else {
				amap[c] = [][2]int{{i, matches[n][0]}}
			}
		}
		i++
	}

	return amap, len(line), i
}

func BuildAntiNodeMap(amap AntennaMap, width int, height int, includeResonance bool) AntiNodeMap {
	antiNodeMap := AntiNodeMap{}

	for _, v := range amap {
		for i := range v {
			for j := i + 1; j < len(v); j++ {

				nodeA, nodeB := v[i], v[j]
				xInc, yInc := nodeB[1]-nodeA[1], nodeB[0]-nodeA[0]

				x, y := nodeA[1], nodeA[0]
				for x >= 0 && x < width && y >= 0 && y < height {
					antiNodeMap[[2]int{y, x}] += 1
					x += xInc
					y += yInc
				}

				x, y = nodeA[1]-xInc, nodeA[0]-yInc
				for x >= 0 && x < width && y >= 0 && y < height {
					antiNodeMap[[2]int{y, x}] += 1
					x -= xInc
					y -= yInc
				}

			}
		}
	}

	return antiNodeMap
}
