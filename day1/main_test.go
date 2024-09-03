package main

import "testing"

func TestParseLine(t *testing.T) {
	type test struct {
		input    string
		expected int
	}

	tests := []test{
		{input: "1abc2", expected: 12},
		{input: "a23fdger4gfdg", expected: 24},
	}
	for _, tc := range tests {
		got := parseLine(tc.input)
		if got != tc.expected {
			t.Fatalf("expected: %d, got: %d", tc.expected, got)
		}
	}
}
