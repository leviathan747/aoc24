package day03

const (
	TOK_OTHER = iota
	TOK_MUL
	TOK_DO
	TOK_DONT
	TOK_COMMA
	TOK_LPAREN
	TOK_RPAREN
	TOK_INTEGER
	TOK_EOF
)

type Lexer struct {
	input  []rune
	cursor int
}

type Token struct {
	tokType int
	text    string
}

func (t Token) String() string {
	str := ""
	switch t.tokType {
	case TOK_OTHER:
		str += "OTHER"
	case TOK_MUL:
		str += "MUL"
	case TOK_DO:
		str += "DO"
	case TOK_DONT:
		str += "DONT"
	case TOK_COMMA:
		str += "COMMA"
	case TOK_LPAREN:
		str += "LPAREN"
	case TOK_RPAREN:
		str += "RPAREN"
	case TOK_INTEGER:
		str += "INTEGER"
	case TOK_EOF:
		str += "EOF"
	}
	if t.tokType == TOK_OTHER || t.tokType == TOK_INTEGER {
		str += "[" + t.text + "]"
	}
	return str
}

func NewLexer(input string) Lexer {
	return Lexer{[]rune(input), 0}
}

func (l *Lexer) NextToken() Token {
	if l.cursor >= len(l.input) {
		return Token{tokType: TOK_EOF, text: ""}
	} else {
		state := 0
		cursorInit := l.cursor
		for {
			if l.cursor < len(l.input) {
				c := l.input[l.cursor]
				l.cursor += 1
				switch state {
				case 0:
					switch c {
					case 'm':
						state = 1
					case ',':
						return Token{tokType: TOK_COMMA, text: string(l.input[cursorInit:l.cursor])}
					case '(':
						return Token{tokType: TOK_LPAREN, text: string(l.input[cursorInit:l.cursor])}
					case ')':
						return Token{tokType: TOK_RPAREN, text: string(l.input[cursorInit:l.cursor])}
					case '0':
						state = 2
					case '1', '2', '3', '4', '5', '6', '7', '8', '9':
						state = 3
					case 'd':
						state = 6
					default:
						state = 5
					}
				case 1:
					switch c {
					case 'u':
						state = 4
					case 'm', ',', '(', ')', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'd':
						l.cursor -= 1
						return Token{tokType: TOK_OTHER, text: string(l.input[cursorInit:l.cursor])}
					default:
						state = 5
					}
				case 2:
					switch c {
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
						state = 3
					default:
						l.cursor -= 1
						return Token{tokType: TOK_INTEGER, text: string(l.input[cursorInit:l.cursor])}
					}
				case 3:
					switch c {
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
						state = 3
					default:
						l.cursor -= 1
						return Token{tokType: TOK_INTEGER, text: string(l.input[cursorInit:l.cursor])}
					}
				case 4:
					switch c {
					case 'l':
						return Token{tokType: TOK_MUL, text: string(l.input[cursorInit:l.cursor])}
					case 'm', ',', '(', ')', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'd':
						l.cursor -= 1
						return Token{tokType: TOK_OTHER, text: string(l.input[cursorInit:l.cursor])}
					default:
						state = 5
					}
				case 5:
					switch c {
					case 'm', ',', '(', ')', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'd':
						l.cursor -= 1
						return Token{tokType: TOK_OTHER, text: string(l.input[cursorInit:l.cursor])}
					default:
						state = 5
					}
				case 6:
					switch c {
					case 'o':
						state = 7
					case 'm', ',', '(', ')', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'd':
						l.cursor -= 1
						return Token{tokType: TOK_DO, text: string(l.input[cursorInit:l.cursor])}
					default:
						state = 5
					}
				case 7:
					switch c {
					case 'n':
						state = 8
					default:
						l.cursor -= 1
						return Token{tokType: TOK_DO, text: string(l.input[cursorInit:l.cursor])}
					}
				case 8:
					switch c {
					case '\'':
						state = 9
					case 'm', ',', '(', ')', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'd':
						l.cursor -= 1
						return Token{tokType: TOK_OTHER, text: string(l.input[cursorInit:l.cursor])}
					default:
						state = 5
					}
				case 9:
					switch c {
					case 't':
						return Token{tokType: TOK_DONT, text: string(l.input[cursorInit:l.cursor])}
					case 'm', ',', '(', ')', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'd':
						l.cursor -= 1
						return Token{tokType: TOK_OTHER, text: string(l.input[cursorInit:l.cursor])}
					default:
						state = 5
					}
				}
			} else {
				switch state {
				case 2, 3:
					return Token{tokType: TOK_INTEGER, text: string(l.input[cursorInit:l.cursor])}
				case 7:
					return Token{tokType: TOK_DO, text: string(l.input[cursorInit:l.cursor])}
				default:
					return Token{tokType: TOK_OTHER, text: string(l.input[cursorInit:l.cursor])}
				}
			}
		}
	}
}
