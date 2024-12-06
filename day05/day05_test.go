// go:build exclude
package day05

import (
	"reflect"
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

	expectedOrderingRules := []OrderingRule{OrderingRule{47, 53},
		OrderingRule{97, 13},
		OrderingRule{97, 61},
		OrderingRule{97, 47},
		OrderingRule{75, 29},
		OrderingRule{61, 13},
		OrderingRule{75, 53},
		OrderingRule{29, 13},
		OrderingRule{97, 29},
		OrderingRule{53, 29},
		OrderingRule{61, 53},
		OrderingRule{97, 53},
		OrderingRule{61, 29},
		OrderingRule{47, 13},
		OrderingRule{75, 47},
		OrderingRule{97, 75},
		OrderingRule{47, 61},
		OrderingRule{75, 61},
		OrderingRule{47, 29},
		OrderingRule{75, 13},
		OrderingRule{53, 13}}

	expectedUpdates := []Update{Update{75, 47, 61, 53, 29},
		Update{97, 61, 53, 29, 13},
		Update{75, 29, 13},
		Update{75, 97, 47, 61, 53},
		Update{61, 13, 29},
		Update{97, 13, 75, 29, 47}}

	orderingRules, updates := ParseInput(testData)

	if !reflect.DeepEqual(orderingRules, expectedOrderingRules) {
		t.Errorf("TestParseInput: Ordering rules do not match:\nexpected: %v\ngot:      %v\n", expectedOrderingRules, orderingRules)
	}

	if !reflect.DeepEqual(updates, expectedUpdates) {
		t.Errorf("TestParseInput: Updates do not match:\nexpected: %v\ngot:      %v\n", expectedUpdates, updates)
	}

}

func TestUpdateIsValid(t *testing.T) {

	orderingRules := []OrderingRule{OrderingRule{47, 53},
		OrderingRule{97, 13},
		OrderingRule{97, 61},
		OrderingRule{97, 47},
		OrderingRule{75, 29},
		OrderingRule{61, 13},
		OrderingRule{75, 53},
		OrderingRule{29, 13},
		OrderingRule{97, 29},
		OrderingRule{53, 29},
		OrderingRule{61, 53},
		OrderingRule{97, 53},
		OrderingRule{61, 29},
		OrderingRule{47, 13},
		OrderingRule{75, 47},
		OrderingRule{97, 75},
		OrderingRule{47, 61},
		OrderingRule{75, 61},
		OrderingRule{47, 29},
		OrderingRule{75, 13},
		OrderingRule{53, 13}}

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
			t.Errorf("TestUpdateIsValid: %v\nexpected: %v\ngot:      %v\n", updates[i], valid, expectedResults[i])
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
