package main

import (
	"reflect"
	"testing"
)

func TestSortedSquaredArray(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedResult []int
	}{
		{
			input:          []int{-4, -1, 0, 3, 10},
			expectedResult: []int{0, 1, 9, 16, 100},
		},
		{
			input:          []int{-7, -3, 2, 3, 11},
			expectedResult: []int{4, 9, 9, 49, 121},
		},
	}

	for _, testCase := range testCases {
		if !reflect.DeepEqual(sortedSquaredArray(testCase.input), testCase.expectedResult) {
			t.Error("Wrong Answer")
		}
	}
}
