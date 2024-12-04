package day03

import (
	"fmt"
	"strconv"
)

const (
	OP_NONE = iota
	OP_MULTIPLY
	OP_ENABLE
	OP_DISABLE
)

type Operation struct {
	operator int
	left     int
	right    int
}

func (o Operation) String() string {
	switch o.operator {
	case OP_MULTIPLY:
		return fmt.Sprintf("%d*%d", o.left, o.right)
	case OP_ENABLE:
		return "enable"
	case OP_DISABLE:
		return "disable"
	default:
		return ""
	}
}

type Parser struct {
	l *Lexer
}

func (p *Parser) Parse() []Operation {
	operations := []Operation{}
	operator := OP_NONE
	left := 0
	right := 0
	nextToken := p.l.NextToken()
	state := 0
	for nextToken.tokType != TOK_EOF {
		switch state {
		case 0:
			switch nextToken.tokType {
			case TOK_MUL:
				operator = OP_MULTIPLY
				state = 1
			case TOK_DO:
				operator = OP_ENABLE
				state = 6
			case TOK_DONT:
				operator = OP_DISABLE
				state = 8
			default:
				state = 0
			}
		case 1:
			switch nextToken.tokType {
			case TOK_LPAREN:
				state = 2
			default:
				state = 0
			}
		case 2:
			switch nextToken.tokType {
			case TOK_INTEGER:
				left, _ = strconv.Atoi(nextToken.text)
				state = 3
			default:
				state = 0
			}
		case 3:
			switch nextToken.tokType {
			case TOK_COMMA:
				state = 4
			default:
				state = 0
			}
		case 4:
			switch nextToken.tokType {
			case TOK_INTEGER:
				right, _ = strconv.Atoi(nextToken.text)
				state = 5
			default:
				state = 0
			}
		case 5:
			switch nextToken.tokType {
			case TOK_RPAREN:
				operations = append(operations, Operation{operator, left, right})
				state = 0
			default:
				state = 0
			}
		case 6:
			switch nextToken.tokType {
			case TOK_LPAREN:
				state = 7
			default:
				state = 0
			}
		case 7:
			switch nextToken.tokType {
			case TOK_RPAREN:
				operations = append(operations, Operation{operator, left, right})
				state = 0
			default:
				state = 0
			}
		case 8:
			switch nextToken.tokType {
			case TOK_LPAREN:
				state = 9
			default:
				state = 0
			}
		case 9:
			switch nextToken.tokType {
			case TOK_RPAREN:
				operations = append(operations, Operation{operator, left, right})
				state = 0
			default:
				state = 0
			}
		}
		nextToken = p.l.NextToken()
	}

	return operations
}
