package day06

import (
	"testing"
)

func TestParseInput(t *testing.T) {

	testData := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	expectedMap := RoomMap{{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, true},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, true, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, true, false},
		{true, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false, false, false}}

	expectedX, expectedY := 4, 6
	expectedXInc, expectedYInc := 0, -1

	theMap, x, y, xInc, yInc := ParseInput(testData)

	if theMap.String() != expectedMap.String() {
		t.Errorf("TestParseInput: Maps do not match:\nexpected:\n%v\ngot:\n%v\n", expectedMap, theMap)
	}

	if x != expectedX || y != expectedY {
		t.Errorf("TestParseInput: Positions do not match: expected (%d, %d) got (%d, %d)\n", expectedX, expectedY, x, y)
	}

	if xInc != expectedXInc || yInc != expectedYInc {
		t.Errorf("TestParseInput: Increments do not match: expected (%d, %d) got (%d, %d)\n", expectedXInc, expectedYInc, xInc, yInc)
	}

}

func TestCountSteps(t *testing.T) {

	theMap := RoomMap{{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, true},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, true, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, true, false},
		{true, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false, false, false}}

	x, y := 4, 6
	xInc, yInc := 0, -1

	expectedSteps := 41

	steps := CountSteps(theMap, x, y, xInc, yInc)

	if steps != expectedSteps {
		t.Errorf("TestCountSteps: expected %d got %d\n", expectedSteps, steps)
	}

}

func TestObstructionLocations(t *testing.T) {

	theMap := RoomMap{{false, false, false, false, true, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, true},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, true, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, true, false},
		{true, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false, false, false}}

	x, y := 4, 6
	xInc, yInc := 0, -1

	expectedLocations := 6

	locations := CountObstructionLocations(theMap, x, y, xInc, yInc)

	if locations != expectedLocations {
		t.Errorf("TestCountLocations: expected %d got %d\n", expectedLocations, locations)
	}

}
