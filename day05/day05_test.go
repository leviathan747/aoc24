// go:build exclude
package day05

import (
	// "fmt"
	"leviathan747/aoc24/input"
	"reflect"
	"slices"
	"testing"
)

func TestParseInput(t *testing.T) {

	testData := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	expectedOrderingRules := OrderingRules{
		29: []int{13},
		47: []int{13, 29, 53, 61},
		53: []int{13, 29},
		61: []int{13, 29, 53},
		75: []int{13, 29, 47, 53, 61},
		97: []int{13, 29, 47, 53, 61, 75}}

	expectedUpdates := []Update{Update{75, 47, 61, 53, 29},
		Update{97, 61, 53, 29, 13},
		Update{75, 29, 13},
		Update{75, 97, 47, 61, 53},
		Update{61, 13, 29},
		Update{97, 13, 75, 29, 47}}

	orderingRules, updates := ParseInput(testData)

	for _, v := range orderingRules {
		slices.Sort(v)
	}

	if !reflect.DeepEqual(orderingRules, expectedOrderingRules) {
		t.Errorf("TestParseInput: Ordering rules do not match:\nexpected: %v\ngot:      %v\n", expectedOrderingRules, orderingRules)
	}

	if !reflect.DeepEqual(updates, expectedUpdates) {
		t.Errorf("TestParseInput: Updates do not match:\nexpected: %v\ngot:      %v\n", expectedUpdates, updates)
	}

}

func TestUpdateIsValid(t *testing.T) {

	orderingRules := OrderingRules{
		29: []int{13},
		47: []int{13, 29, 53, 61},
		53: []int{13, 29},
		61: []int{13, 29, 53},
		75: []int{13, 29, 47, 53, 61},
		97: []int{13, 29, 47, 53, 61, 75}}

	updates := []Update{Update{75, 47, 61, 53, 29},
		Update{97, 61, 53, 29, 13},
		Update{75, 29, 13},
		Update{75, 97, 47, 61, 53},
		Update{61, 13, 29},
		Update{97, 13, 75, 29, 47}}

	expectedResults := []bool{true, true, true, false, false, false}

	for i := 0; i < len(updates); i++ {

		valid := UpdateIsValid(updates[i], orderingRules)

		if valid != expectedResults[i] {
			t.Errorf("TestUpdateIsValid: %v\nexpected: %v\ngot:      %v\n", updates[i], expectedResults[i], valid)
		}

	}

}

func TestSumMiddles(t *testing.T) {

	updates := []Update{Update{75, 47, 61, 53, 29},
		Update{97, 61, 53, 29, 13},
		Update{75, 29, 13}}

	expectedSum := 143

	sum := SumMiddles(updates)

	if sum != expectedSum {
		t.Errorf("TestSumMiddles: Incorrect sum: expected: %v\ngot:      %v\n", expectedSum, sum)
	}

}

func TestFixInvalids(t *testing.T) {

	orderingRules := OrderingRules{
		29: []int{13},
		47: []int{13, 29, 53, 61},
		53: []int{13, 29},
		61: []int{13, 29, 53},
		75: []int{13, 29, 47, 53, 61},
		97: []int{13, 29, 47, 53, 61, 75}}

	updates := []Update{Update{75, 97, 47, 61, 53},
		Update{61, 13, 29},
		Update{97, 13, 75, 29, 47}}

	expectedSum := 123

	FixInvalidUpdates(updates, orderingRules)

	sum := SumMiddles(updates)

	if sum != expectedSum {
		t.Errorf("TestFixInvalids: Incorrect sum: expected %d, got %d\n", expectedSum, sum)
	}

}

func TestPart1(t *testing.T) {
	t.Skip()
	input := input.GetInput("./day05_input.txt")
	orderingRules, updates := ParseInput(input)
	validUpdates := []Update{}
	for i := 0; i < len(updates); i++ {
		// fmt.Printf("Checking (%d/%d): %v\n", i+1, len(updates), updates[i])
		if UpdateIsValid(updates[i], orderingRules) {
			validUpdates = append(validUpdates, updates[i])
		}
	}
	sum := SumMiddles(validUpdates)
	expectedSum := 5087
	if sum != expectedSum {
		t.Errorf("TestPart1: Incorrect sum: expected %d, got %d\n", expectedSum, sum)
	}
}
