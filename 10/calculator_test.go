// Package calculator test module
package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 2},
		{a: 2, b: 2, want: 4},
		{a: 6, b: 7, want: 13},
		{a: -1, b: -1, want: -2},
		{a: -1, b: 1, want: 0},
		{a: 1, b: 1, want: 2},
	}
	for _, testCase := range testCases {
		got := calculator.Add(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Add(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}
func TestSubtract(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 5, b: 4, want: 1},
		{a: 10, b: 11, want: -1},
		{a: 0, b: 0, want: 0},
		{a: 100, b: 40, want: 60},
		{a: -1, b: -1, want: 0},
	}
	for _, testCase := range testCases {
		result := calculator.Subtract(testCase.a, testCase.b)
		if testCase.want != result {
			t.Errorf("Subtract(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, result)
		}
	}
}
func TestMultiply(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 5, b: 5, want: 25},
		{a: 4, b: 5, want: 20},
		{a: 1, b: 1, want: 1},
		{a: 2, b: 10, want: 20},
		{a: 50, b: 2, want: 100},
	}
	for _, testCase := range testCases {
		result := calculator.Multiply(testCase.a, testCase.b)
		if testCase.want != result {
			t.Errorf("Multiply(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, result)
		}
	}
}
func TestDivide(t *testing.T) {
	testCases := []struct {
		a, b        int
		want        int
		errExpected bool
	}{
		{a: 10, b: 5, want: 2, errExpected: false},
		{a: 100, b: 2, want: 50, errExpected: false},
		{a: 25, b: 5, want: 5, errExpected: false},
		{a: 10, b: 1, want: 10, errExpected: false},
		{a: 15, b: 3, want: 5, errExpected: false},
	}
	for _, testCase := range testCases {
		result, err := calculator.Divide(testCase.a, testCase.b)
		// test if we got an error
		if err != nil {
			// test if we got an expected error
			if !testCase.errExpected {
				t.Errorf("Divide(%d, %d): want %d, got %d, error %d", testCase.a, testCase.b, testCase.want, result, err)
			}
		} else {
			// test if we expected an error and go tnilo
			if testCase.errExpected {
				t.Errorf("Divide(%d, %d): want error got nil", testCase.a, testCase.b)
			}
			// test data value - want equals result
			if testCase.want != result {
				t.Errorf("Divide(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, result)
			}
		}
	}
}
