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

func TestDecodeDigits(t *testing.T) {
	type test struct {
		input    string
		expected string
	}

	tests := []test{
		{input: "two1nine", expected: "2o19e"},
		{input: "eighttwothree", expected: "8t2o3e"},
		{input: "zoneeight234", expected: "z1e8t234"},
		{input: "esevensevenseven", expected: "e7n7n7n"},
		{input: "xtwone3four", expected: "x21e34"},
		{input: "oneight", expected: "18t"},
	}
	for _, tc := range tests {
		got := decodeDigits(tc.input)
		if got != tc.expected {
			t.Fatalf("expected: %s, got: %s", tc.expected, got)
		}
	}
}
