package day03

import "strconv"

const (
	OP_NONE = iota
	OP_MULTIPLY
)

type Operation struct {
	operator int
	left     int
	right    int
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
		}
		nextToken = p.l.NextToken()
	}

	return operations
}
