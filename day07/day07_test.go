package day07

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {

	testData := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	expectedCalibrations := []Calibration{
		{190, []int{10, 19}},
		{3267, []int{81, 40, 27}},
		{83, []int{17, 5}},
		{156, []int{15, 6}},
		{7290, []int{6, 8, 6, 15}},
		{161011, []int{16, 10, 13}},
		{192, []int{17, 8, 14}},
		{21037, []int{9, 7, 18, 13}},
		{292, []int{11, 6, 16, 20}}}

	calibrations := ParseInput(testData)

	if !reflect.DeepEqual(calibrations, expectedCalibrations) {

		t.Errorf("TestParseInput:\nexpected: %v\ngot: %v\n", expectedCalibrations, calibrations)

	}

}

func TestCalibrationIsValid(t *testing.T) {

	calibrations := []Calibration{
		{190, []int{10, 19}},
		{3267, []int{81, 40, 27}},
		{83, []int{17, 5}},
		{156, []int{15, 6}},
		{7290, []int{6, 8, 6, 15}},
		{161011, []int{16, 10, 13}},
		{192, []int{17, 8, 14}},
		{21037, []int{9, 7, 18, 13}},
		{292, []int{11, 6, 16, 20}}}

	expectedResults := []bool{true, true, false, false, false, false, false, false, true}

	for i := range calibrations {

		valid := CalibrationIsValid(calibrations[i])

		if valid != expectedResults[i] {

			t.Errorf("TestCalibrationIsValid: %v expected %v got %v\n", calibrations[i], expectedResults[i], valid)

		}

	}

}
