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
	answer := ComputeResult(operations)
	fmt.Println(answer)
}

func ComputeResult(operations []Operation) int {

	sum := 0

	for i := 0; i < len(operations); i++ {
		operation := operations[i]
		switch operation.operator {
		case OP_MULTIPLY:
			sum += operation.left * operation.right
		}
	}

	return sum

}
