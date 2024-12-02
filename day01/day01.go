package day01

import (
	"bufio"
	"fmt"
	"leviathan747/aoc24/input"
	"slices"
	"strconv"
	"strings"
)

func Day01() {
	input := input.GetInput("./day01/day01_input.txt")
	list1, list2 := ParseInput(input)
	dist := ComputeDistance(list1, list2)
	fmt.Println(dist)
}

func ParseInput(input string) ([]int, []int) {
	list1 := []int{}
	list2 := []int{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		num1, _ := strconv.Atoi(fields[0])
		num2, _ := strconv.Atoi(fields[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scan error occurred: %v\n", err)
	}
	return list1, list2
}

func ComputeDistance(list1, list2 []int) int {
	// note: this assumes the input lists are equal in length

	// sort the lists
	slices.Sort(list1)
	slices.Sort(list2)

	// loop through and add the distances
	dist := 0
	for i := 0; i < len(list1); i++ {
		if list1[i] > list2[i] {
			dist += list1[i] - list2[i]
		} else {
			dist += list2[i] - list1[i]
		}
	}

	return dist

}
