package calculator_test

import (
	"calculator"
	"testing"
)

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
		a, b int
		want int
	}{
		{a: 10, b: 5, want: 2},
		{a: 100, b: 2, want: 50},
		{a: 25, b: 5, want: 5},
		{a: 10, b: 1, want: 10},
		{a: 15, b: 3, want: 5},
	}
	for _, testCase := range testCases {
		result := calculator.Division(testCase.a, testCase.b)
		if testCase.want != result {
			t.Errorf("Divide(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, result)
		}
	}
}