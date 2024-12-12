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
	antiNodeMap := BuildAntiNodeMap(amap, width, height)
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

func BuildAntiNodeMap(amap AntennaMap, width int, height int) AntiNodeMap {
	antiNodeMap := AntiNodeMap{}

	for _, v := range amap {
		for i := range v {
			for j := i + 1; j < len(v); j++ {
				nodeA, nodeB := v[i], v[j]
				x1, y1 := 2*nodeA[1]-nodeB[1], 2*nodeA[0]-nodeB[0]
				if x1 >= 0 && x1 < width && y1 >= 0 && y1 < height {
					antiNodeMap[[2]int{y1, x1}] += 1
				}
				x2, y2 := 2*nodeB[1]-nodeA[1], 2*nodeB[0]-nodeA[0]
				if x2 >= 0 && x2 < width && y2 >= 0 && y2 < height {
					antiNodeMap[[2]int{y2, x2}] += 1
				}
			}
		}

	}

	return antiNodeMap
}
