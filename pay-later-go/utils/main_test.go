package utils

import "testing"

func TestFormatFloat(t *testing.T) {
	testCases := []struct {
		num      float64
		expected string
	}{
		{num: 2.50, expected: "2.5"},
		{num: 2.645, expected: "2.645"},
		{num: 2.0, expected: "2"},
		{num: 0.120, expected: "0.12"},
		{num: 0.123456, expected: "0.123456"},
	}

	for _, tc := range testCases {
		got := FormatFloat(tc.num)
		if got != tc.expected {
			t.Fatalf("expected: %v, got: %v", tc.expected, got)
		}
	}
}
