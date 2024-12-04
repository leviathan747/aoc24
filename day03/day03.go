package day03

import (
	"fmt"
	"leviathan747/aoc24/input"
)

func Day03() {
	input := input.GetInput("./day03/day03_input.txt")
	l := NewLexer(input)
	p := Parser{&l}
	operations := p.Parse()
	answer := ComputeResult(operations, false)
	fmt.Println(answer)
	answer = ComputeResult(operations, true)
	fmt.Println(answer)
}

func ComputeResult(operations []Operation, conditionals bool) int {

	sum := 0
	enabled := true

	for i := 0; i < len(operations); i++ {
		operation := operations[i]
		switch operation.operator {
		case OP_MULTIPLY:
			if enabled || !conditionals {
				sum += operation.left * operation.right
			}
		case OP_ENABLE:
			enabled = true
		case OP_DISABLE:
			enabled = false
		}
	}

	return sum

}
