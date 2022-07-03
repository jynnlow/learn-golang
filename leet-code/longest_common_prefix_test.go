package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		input          []string
		expectedResult string
	}{
		{
			input:          []string{"flower", "flow", "flight"},
			expectedResult: "fl",
		},
		{
			input:          []string{"dog", "racecar", "car"},
			expectedResult: "",
		},
	}

	for _, testCase := range testCases {
		result := longestCommonPrefix(testCase.input)
		if result != testCase.expectedResult {
			t.Error("wrong answer")
		}
	}
}
