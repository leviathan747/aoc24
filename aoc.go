package main

import (
	"fmt"
	"leviathan747/aoc24/day01"
	"leviathan747/aoc24/day02"
	"leviathan747/aoc24/day03"
	"leviathan747/aoc24/day04"
	"leviathan747/aoc24/day05"
	"leviathan747/aoc24/day06"
	"leviathan747/aoc24/day07"
	"leviathan747/aoc24/day08"
	"os"
)

func main() {
	// Ensure there's at least one command line argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a function name as a commandline argument.")
		return
	}

	// Get the function name from the first argument
	funcName := os.Args[1]

	// Map function names to their corresponding functions
	functions := map[string]func(){
		"day01": day01.Day01,
		"day02": day02.Day02,
		"day03": day03.Day03,
		"day04": day04.Day04,
		"day05": day05.Day05,
		"day06": day06.Day06,
		"day07": day07.Day07,
		"day08": day08.Day08,
	}

	// Check if the function exists and call it
	if function, exists := functions[funcName]; exists {
		function()
	} else {
		fmt.Printf("Function '%s' not found.\n", funcName)
	}
}
