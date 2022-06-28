package main

import "testing"

func TestSearchInputPosition(t *testing.T) {
	testCases := []struct {
		input          []int
		target         int
		expectedResult int
	}{
		{
			input:          []int{1, 3, 5, 6},
			target:         5,
			expectedResult: 2,
		}, {
			input:          []int{1, 3, 5, 6},
			target:         2,
			expectedResult: 1,
		}, {
			input:          []int{1, 3, 5, 6},
			target:         7,
			expectedResult: 4,
		},
	}

	for _, testCase := range testCases {
		if searchInputPosition(testCase.input, testCase.target) != testCase.expectedResult {
			t.Error("wrong answer")
		}
	}
}
