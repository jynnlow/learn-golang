package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		input          []int
		target         int
		expectedResult []int
	}{
		{
			input:          []int{2, 7, 11, 15},
			target:         9,
			expectedResult: []int{1, 2},
		},
		{
			input:          []int{2, 3, 4},
			target:         6,
			expectedResult: []int{1, 3},
		},
		{
			input:          []int{-1, 0},
			target:         -1,
			expectedResult: []int{1, 2},
		},
	}

	for _, testCase := range testCases {
		result := twoSum(testCase.input, testCase.target)
		if !reflect.DeepEqual(result, testCase.expectedResult) {
			t.Error("wrong answer")
		}
	}
}
