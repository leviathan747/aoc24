package day02

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {

	testData := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	expectedReports := [][]int{{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9}}

	reports := ParseInput(testData)

	if !reflect.DeepEqual(reports, expectedReports) {
		t.Errorf("TestParseInput: Input lists do not match\n")
	}

}

func TestNumSafeReports(t *testing.T) {

	reports := [][]int{{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9}}

	numSafe := NumSafeReports(reports)

	expectedSafe := 2

	if numSafe != expectedSafe {
		t.Errorf("TestNumSafeReports: expected %d got %d\n", expectedSafe, numSafe)
	}

}
