package main

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	testCases := []struct {
		input          []int
		steps          int
		expectedResult []int
	}{
		{
			input:          []int{1, 2, 3, 4, 5, 6, 7},
			steps:          3,
			expectedResult: []int{5, 6, 7, 1, 2, 3, 4},
		}, {
			input:          []int{-1, -100, 3, 99},
			steps:          2,
			expectedResult: []int{3, 99, -1, -100},
		},
	}

	for _, testCase := range testCases {
		if !reflect.DeepEqual(rotateArray(testCase.input, testCase.steps), testCase.expectedResult) {
			t.Error("wrong answer")
		}
	}
}
