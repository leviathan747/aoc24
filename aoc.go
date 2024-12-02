package main

import (
	"fmt"
	"leviathan747/aoc24/day01"
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
	}

	// Check if the function exists and call it
	if function, exists := functions[funcName]; exists {
		function()
	} else {
		fmt.Printf("Function '%s' not found.\n", funcName)
	}
}
