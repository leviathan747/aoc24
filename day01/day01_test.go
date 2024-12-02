package day01

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {

	testData := `3   4
4   3
2   5
1   3
3   9
3   3`

	list1, list2 := ParseInput(testData)

	expectedList1 := []int{3, 4, 2, 1, 3, 3}
	expectedList2 := []int{4, 3, 5, 3, 9, 3}

	if !reflect.DeepEqual(list1, expectedList1) || !reflect.DeepEqual(list2, expectedList2) {
		t.Errorf("ParseInput: Input lists do not match\n")
	}

}

func TestComputeDistance(t *testing.T) {

	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	dist := ComputeDistance(list1, list2)
	expectedResult := 11
	if dist != expectedResult {
		t.Errorf("ComputeDistance: got %d, expected %d\n", dist, expectedResult)
	}

}
