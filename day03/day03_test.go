package day03

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLexInput(t *testing.T) {

	testStrings := []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
		"othermul3othermu2",
		"othermul3otherm2"}

	expectedTokLists := []string{"OTHER[x] MUL LPAREN INTEGER[2] COMMA INTEGER[4] RPAREN OTHER[%&] MUL OTHER[[] INTEGER[3] COMMA INTEGER[7] OTHER[]!@^do_not_] MUL LPAREN INTEGER[5] COMMA INTEGER[5] RPAREN OTHER[+] MUL LPAREN INTEGER[32] COMMA INTEGER[64] OTHER[]then] LPAREN MUL LPAREN INTEGER[11] COMMA INTEGER[8] RPAREN MUL LPAREN INTEGER[8] COMMA INTEGER[5] RPAREN RPAREN",
		"OTHER[other] MUL INTEGER[3] OTHER[other] OTHER[mu] INTEGER[2]", // Lexer does not look ahead to detect single OTHER token
		"OTHER[other] MUL INTEGER[3] OTHER[other] OTHER[m] INTEGER[2]"}

	for i := 0; i < len(testStrings); i++ {
		testData := testStrings[i]
		expectedTokList := expectedTokLists[i]

		l := NewLexer(testData)

		tokList := ""
		sep := ""
		nextToken := l.NextToken()
		for i := 0; i < 100 && nextToken.tokType != TOK_EOF; i++ {
			tokList = fmt.Sprintf("%s%s%s", tokList, sep, nextToken)
			sep = " "
			nextToken = l.NextToken()
		}

		if tokList != expectedTokList {
			t.Errorf("TestLexInput:\nexpected: '%s'\ngot:      '%s'\n", expectedTokList, tokList)
		}

	}

}

func TestParseInput(t *testing.T) {

	testData := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	expectedOperations := []Operation{
		Operation{OP_MULTIPLY, 2, 4},
		Operation{OP_MULTIPLY, 5, 5},
		Operation{OP_MULTIPLY, 11, 8},
		Operation{OP_MULTIPLY, 8, 5}}

	l := NewLexer(testData)
	p := Parser{&l}
	operations := p.Parse()

	if !reflect.DeepEqual(operations, expectedOperations) {
		t.Errorf("TestParseInput: Operation lists do not match\n")
	}

}

func TestComputeResult(t *testing.T) {

	operations := []Operation{
		Operation{OP_MULTIPLY, 2, 4},
		Operation{OP_MULTIPLY, 5, 5},
		Operation{OP_MULTIPLY, 11, 8},
		Operation{OP_MULTIPLY, 8, 5}}

	expectedSum := 161

	sum := ComputeResult(operations)

	if sum != expectedSum {
		t.Errorf("TestComputeResult: expected %d, got %d\n", expectedSum, sum)
	}

}
