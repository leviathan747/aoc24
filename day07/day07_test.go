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

		valid := CalibrationIsValid(calibrations[i], []Operation{Multiply, Add})

		if valid != expectedResults[i] {

			t.Errorf("TestCalibrationIsValid: %v expected %v got %v\n", calibrations[i], expectedResults[i], valid)

		}

	}

}

func TestCalibrationIsValidWithConcat(t *testing.T) {

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

	expectedResults := []bool{true, true, false, true, true, false, true, false, true}

	for i := range calibrations {

		valid := CalibrationIsValid(calibrations[i], []Operation{Multiply, Add, Concat})

		if valid != expectedResults[i] {

			t.Errorf("TestCalibrationIsValidWithConcat: %v expected %v got %v\n", calibrations[i], expectedResults[i], valid)

		}

	}

}

func TestPowLog(t *testing.T) {

	x := Pow(10, 0)
	y := 1
	if x != y {
		t.Errorf("TestPowLog: expected %d got %d\n", y, x)
	}

	x = Pow(10, 1)
	y = 10
	if x != y {
		t.Errorf("TestPowLog: expected %d got %d\n", y, x)
	}

	x = Pow(10, 2)
	y = 100
	if x != y {
		t.Errorf("TestPowLog: expected %d got %d\n", y, x)
	}

	x = LogBase(1, 10)
	y = 0
	if x != y {
		t.Errorf("TestPowLog: expected %d got %d\n", y, x)
	}

	x = LogBase(9, 10)
	y = 0
	if x != y {
		t.Errorf("TestPowLog: expected %d got %d\n", y, x)
	}

	x = LogBase(10, 10)
	y = 1
	if x != y {
		t.Errorf("TestPowLog: expected %d got %d\n", y, x)
	}

	x = LogBase(99, 10)
	y = 1
	if x != y {
		t.Errorf("TestPowLog: expected %d got %d\n", y, x)
	}

	x = LogBase(100, 10)
	y = 2
	if x != y {
		t.Errorf("TestPowLog: expected %d got %d\n", y, x)
	}

	x = LogBase(999, 10)
	y = 2
	if x != y {
		t.Errorf("TestPowLog: expected %d got %d\n", y, x)
	}

}
