package day04

import (
	"bufio"
	"fmt"
	"leviathan747/aoc24/input"
	"strings"
)

func Day04() {
	input := input.GetInput("./day04/day04_input.txt")
	matrix := ParseInput(input)
	count := FindXmas(matrix)
	fmt.Println(count)
}

func ParseInput(input string) [][]rune {
	matrix := [][]rune{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scan error occurred: %v\n", err)
	}
	return matrix
}

func FindXmas(matrix [][]rune) int {
	// create one big search string and then count the occurences of XMAS
	sb := strings.Builder{}

	// add each row forwards
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			sb.WriteRune(matrix[i][j])
		}
		sb.WriteRune(' ')
	}
	sb.WriteRune('\n')

	// add each row backwards
	for i := 0; i < len(matrix); i++ {
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			sb.WriteRune(matrix[i][j])
		}
		sb.WriteRune(' ')
	}
	sb.WriteRune('\n')

	// add each column top down
	if len(matrix) > 0 {
		for i := 0; i < len(matrix[0]); i++ {
			for j := 0; j < len(matrix); j++ {
				sb.WriteRune(matrix[j][i])
			}
			sb.WriteRune(' ')
		}
	}
	sb.WriteRune('\n')

	// add each column bottom up
	if len(matrix) > 0 {
		for i := 0; i < len(matrix[0]); i++ {
			for j := len(matrix) - 1; j >= 0; j-- {
				sb.WriteRune(matrix[j][i])
			}
			sb.WriteRune(' ')
		}
	}
	sb.WriteRune('\n')

	// add southeasterly diagonals
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 || i == 0 { // only start on the far left or top edges
				for x, y := j, i; x < len(matrix[i]) && y < len(matrix); x, y = x+1, y+1 {
					sb.WriteRune(matrix[y][x])
				}
				sb.WriteRune(' ')
			}
		}
	}
	sb.WriteRune('\n')

	// add northeasterly diagonals
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 || i == len(matrix)-1 { // only start on the far left or bottom edges
				for x, y := j, i; x < len(matrix[i]) && y >= 0; x, y = x+1, y-1 {
					sb.WriteRune(matrix[y][x])
				}
				sb.WriteRune(' ')
			}
		}
	}
	sb.WriteRune('\n')

	// add southwesterly diagonals
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == len(matrix[i])-1 || i == 0 { // only start on the far right or top edges
				for x, y := j, i; x >= 0 && y < len(matrix); x, y = x-1, y+1 {
					sb.WriteRune(matrix[y][x])
				}
				sb.WriteRune(' ')
			}
		}
	}
	sb.WriteRune('\n')

	// add northwesterly diagonals
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == len(matrix[i])-1 || i == len(matrix)-1 { // only start on the far right or bottom edges
				for x, y := j, i; x >= 0 && y >= 0; x, y = x-1, y-1 {
					sb.WriteRune(matrix[y][x])
				}
				sb.WriteRune(' ')
			}
		}
	}
	sb.WriteRune('\n')

	// fmt.Println(sb.String())

	return strings.Count(sb.String(), "XMAS")
}
