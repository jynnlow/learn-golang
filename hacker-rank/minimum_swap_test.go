package main

import (
	"testing"
)

type minimumSwapTestStruct struct {
	arr []int32
	answer int
}

func Test_minimumSwap(t *testing.T){
	testMinimumSwapArray := []minimumSwapTestStruct{
		{
			arr: []int32{
				7, 1, 3, 2, 4, 5, 6,
			},
			answer: 5,
		},
		{
			arr: []int32{
				1, 3, 5, 2, 4, 6, 7,
			},
			answer: 3,
		},
		{
			arr: []int32{
				2, 3, 4, 1, 5,
			},
			answer: 3,
		},
		{
			arr: []int32{
				4, 3, 1, 2,
			},
			answer: 3,
		},
	}

	for _, totalSwap := range testMinimumSwapArray{
		if minimumSwapUsingMap(totalSwap.arr) != int32(totalSwap.answer) {
			t.Error("Wrong Answer")
		}
	}
}
