package main

import (
	"reflect"
	"testing"
)

func TestIntersectionOfTwoArrays(t *testing.T) {
	testCases := []struct {
		input          []int
		input2         []int
		expectedResult []int
	}{
		{
			input:          []int{1, 2, 2, 1},
			input2:         []int{2, 2},
			expectedResult: []int{2, 2},
		},
		{
			input:          []int{4, 9, 5},
			input2:         []int{9, 4, 9, 8, 4},
			expectedResult: []int{9, 4},
		},
	}

	for _, testCase := range testCases {
		result := intersectionOfTwoArrays(testCase.input, testCase.input2)
		if !(reflect.DeepEqual(result, testCase.expectedResult)) {
			t.Error("wrong answer")
		}
	}
}
