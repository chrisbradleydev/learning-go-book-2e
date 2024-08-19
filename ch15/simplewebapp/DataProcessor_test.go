package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func Test_parser(t *testing.T) {
	data1 := strings.Join([]string{"CALC_1", "+", "2", "3"}, "\n")
	data2 := strings.Join([]string{"CALC_2", "-", "5", "1"}, "\n")
	testCases := []struct {
		data     []byte
		expected Input
	}{
		{[]byte(data1), Input{Id: "CALC_1", Op: "+", Val1: 2, Val2: 3}},
		{[]byte(data2), Input{Id: "CALC_2", Op: "-", Val1: 5, Val2: 1}},
	}

	for _, tc := range testCases {
		result, _ := parser(tc.data)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Parser function output is incorrect. Expected: %v, but got: %v", tc.expected, result)
		}
	}
}

func TestDataProcessor(t *testing.T) {
	data1 := strings.Join([]string{"CALC_1", "*", "3", "2"}, "\n")
	data2 := strings.Join([]string{"CALC_2", "/", "21", "7"}, "\n")
	testCases := []struct {
		data   []byte
		result Result
	}{
		{[]byte(data1), Result{Id: "CALC_1", Value: 6}},
		{[]byte(data2), Result{Id: "CALC_2", Value: 3}},
	}

	in := make(chan []byte)
	out := make(chan Result)

	go func() {
		for _, tc := range testCases {
			in <- tc.data
		}
		close(in)
	}()

	go DataProcessor(in, out)

	var results []Result
	for r := range out {
		results = append(results, r)
	}

	expectedResults := []Result{
		{Id: "CALC_1", Value: 6},
		{Id: "CALC_2", Value: 3},
	}

	if !reflect.DeepEqual(results, expectedResults) {
		t.Errorf("DataProcessor function output is incorrect. Expected: %v, but got: %v", expectedResults, results)
	}
}

func TestWriteData(t *testing.T) {
	results := []Result{
		{Id: "CALC_1", Value: 100},
		{Id: "CALC_2", Value: 200},
	}

	in := make(chan Result)
	w := &bytes.Buffer{}

	go func() {
		for _, r := range results {
			in <- r
		}
		close(in)
	}()

	WriteData(in, w)

	expectedOutput := "CALC_1:100\nCALC_2:200\n"
	if !reflect.DeepEqual(w.String(), expectedOutput) {
		t.Errorf("incorrect result: expected: %s, got: %s", expectedOutput, w.String())
	}
}
