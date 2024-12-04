package day03

import (
	"fmt"
	"testing"
)

func TestLexInput(t *testing.T) {

	testStrings := []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
		"othermul3othermu2",
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}

	expectedTokLists := []string{"OTHER[x] MUL LPAREN INTEGER[2] COMMA INTEGER[4] RPAREN OTHER[%&] MUL OTHER[[] INTEGER[3] COMMA INTEGER[7] OTHER[]!@^] DO OTHER[_not_] MUL LPAREN INTEGER[5] COMMA INTEGER[5] RPAREN OTHER[+] MUL LPAREN INTEGER[32] COMMA INTEGER[64] OTHER[]then] LPAREN MUL LPAREN INTEGER[11] COMMA INTEGER[8] RPAREN MUL LPAREN INTEGER[8] COMMA INTEGER[5] RPAREN RPAREN",
		"OTHER[other] MUL INTEGER[3] OTHER[other] OTHER[mu] INTEGER[2]", // Lexer does not look ahead to detect single OTHER token
		"OTHER[x] MUL LPAREN INTEGER[2] COMMA INTEGER[4] RPAREN OTHER[&] MUL OTHER[[] INTEGER[3] COMMA INTEGER[7] OTHER[]!^] DONT LPAREN RPAREN OTHER[_] MUL LPAREN INTEGER[5] COMMA INTEGER[5] RPAREN OTHER[+] MUL LPAREN INTEGER[32] COMMA INTEGER[64] OTHER[]] LPAREN MUL LPAREN INTEGER[11] COMMA INTEGER[8] RPAREN OTHER[un] DO LPAREN RPAREN OTHER[?] MUL LPAREN INTEGER[8] COMMA INTEGER[5] RPAREN RPAREN"}

	for i := 0; i < len(testStrings); i++ {
		testData := testStrings[i]
		expectedTokList := expectedTokLists[i]

		l := NewLexer(testData)

		tokList := ""
		sep := ""
		nextToken := l.NextToken()
		for nextToken.tokType != TOK_EOF {
			tokList = fmt.Sprintf("%s%s%s", tokList, sep, nextToken)
			sep = " "
			nextToken = l.NextToken()
		}

		if tokList != expectedTokList {
			t.Errorf("TestLexInput: %s\nexpected: '%s'\ngot:      '%s'\n", testData, expectedTokList, tokList)
		}

	}

}

func TestParseInput(t *testing.T) {

	testStrings := []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}

	expectedOperationSets := []string{
		"2*4, 5*5, 11*8, 8*5",
		"2*4, disable, 5*5, 11*8, enable, 8*5"}

	for i := 0; i < len(testStrings); i++ {
		testData := testStrings[i]
		expectedOperations := expectedOperationSets[i]

		l := NewLexer(testData)
		p := Parser{&l}
		operations := p.Parse()

		opList := ""
		sep := ""
		for i := 0; i < len(operations); i++ {
			opList = fmt.Sprintf("%s%s%s", opList, sep, operations[i])
			sep = ", "
		}

		if opList != expectedOperations {
			t.Errorf("TestParseInput: Operation lists do not match\nexpected: '%s'\ngot:      '%s'\n", expectedOperations, opList)
		}

	}

}

func TestComputeResult(t *testing.T) {

	operationSets := [][]Operation{{
		Operation{OP_MULTIPLY, 2, 4},
		Operation{OP_MULTIPLY, 5, 5},
		Operation{OP_MULTIPLY, 11, 8},
		Operation{OP_MULTIPLY, 8, 5}}, {
		Operation{OP_MULTIPLY, 2, 4},
		Operation{OP_DISABLE, 0, 0},
		Operation{OP_MULTIPLY, 5, 5},
		Operation{OP_MULTIPLY, 11, 8},
		Operation{OP_ENABLE, 0, 0},
		Operation{OP_MULTIPLY, 8, 5}}}

	expectedSums := []int{161, 48}

	for i := 0; i < len(operationSets); i++ {
		operations := operationSets[i]
		expectedSum := expectedSums[i]

		sum := ComputeResult(operations, true)

		if sum != expectedSum {
			t.Errorf("TestComputeResult: expected %d, got %d\n", expectedSum, sum)
		}

	}

}
