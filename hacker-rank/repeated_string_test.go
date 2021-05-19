package main

import (
	"strconv"
	"testing"
)

type repeatedStringStruct struct {
	substring string
	maxlength int
	answer int
}

func Test_findA(t *testing.T) {
	testRepeatedStringArr := []repeatedStringStruct{
		{
			substring: "gfcaaaecbg",
			maxlength: 547602,
			answer: 164280,
		},
	}

	for index, testRepeatedString := range testRepeatedStringArr {
		if findA(testRepeatedString.substring, testRepeatedString.maxlength) != testRepeatedString.answer {
			t.Error(testRepeatedString.substring + " is wrong at index " + strconv.Itoa(index))
		}
	}
}