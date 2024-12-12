package day08

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {

	testData := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

	expectedAMap := AntennaMap{'0': [][2]int{{1, 8}, {2, 5}, {3, 7}, {4, 4}},
		'A': [][2]int{{5, 6}, {8, 8}, {9, 9}}}

	expectedWidth, expectedHeight := 12, 12

	amap, width, height := ParseInput(testData)

	if !reflect.DeepEqual(amap, expectedAMap) {
		t.Errorf("TestParseInput:\nexpected: %v\ngot: %v\n", expectedAMap, amap)
	}

	if width != expectedWidth {
		t.Errorf("TestParseInput: Wrong width: expected %d got %d\n", expectedWidth, width)
	}

	if height != expectedHeight {
		t.Errorf("TestParseInput: Wrong height: expected %d got %d\n", expectedHeight, height)
	}

}

func TestBuildAntiNodeMap(t *testing.T) {

	amap := AntennaMap{'0': [][2]int{{1, 8}, {2, 5}, {3, 7}, {4, 4}},
		'A': [][2]int{{5, 6}, {8, 8}, {9, 9}}}

	expectedAntiNodeMap := AntiNodeMap{{0, 6}: 1,
		{0, 11}:  1,
		{1, 3}:   2,
		{2, 4}:   1,
		{2, 10}:  1,
		{3, 2}:   1,
		{4, 9}:   1,
		{5, 1}:   1,
		{5, 6}:   1,
		{6, 3}:   1,
		{7, 0}:   1,
		{7, 7}:   1,
		{10, 10}: 1,
		{11, 10}: 1}

	antiNodeMap := BuildAntiNodeMap(amap, 12, 12)

	if !reflect.DeepEqual(antiNodeMap, expectedAntiNodeMap) {
		t.Errorf("TestBuildAntiNodeMap:\nexpected: %v\ngot:      %v\n", expectedAntiNodeMap, antiNodeMap)
	}

}
