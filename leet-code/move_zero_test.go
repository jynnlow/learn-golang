package main

import (
	"reflect"
	"testing"
)

func TestMoveZero(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedResult []int
	}{
		{
			input:          []int{0, 1, 0, 3, 12},
			expectedResult: []int{1, 3, 12, 0, 0},
		},
		{
			input:          []int{0},
			expectedResult: []int{0},
		},
	}

	for _, testCase := range testCases {
		if !reflect.DeepEqual(moveZero(testCase.input), testCase.expectedResult) {
			t.Error("wrong answer")
		}
	}
}
