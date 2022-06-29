package main

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		input          []byte
		expectedResult []byte
	}{
		{
			input:          []byte("hello"),
			expectedResult: []byte("olleh"),
		},
		{
			input:          []byte("Hannah"),
			expectedResult: []byte("hannaH"),
		},
	}

	for _, testCase := range testCases {
		if !reflect.DeepEqual(reverseString(testCase.input), testCase.expectedResult) {
			t.Error("wrong answer")
		}
	}
}
