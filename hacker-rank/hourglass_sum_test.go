package main

import "testing"

type hourglass struct{
	arr [][]int32
	answer int32
}

func Test_hourglassSum(t *testing.T) {
	testHourglassSumArr := []hourglass{
		{
			arr: [][]int32{
				{1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 0},
				{1, 1, 1, 0, 0, 0},
				{0, 0, 2, 4, 4, 0},
				{0, 0, 0, 2, 0, 0},
				{0, 0, 1, 2, 4, 0},
			},
			answer: 19,
		},
	}

	for _, testHourglassSum := range testHourglassSumArr {
		if hourglassSum(testHourglassSum.arr) != testHourglassSum.answer {
			t.Error("Wrong Answer")
		}
	}

	//for i := 0; i < len(testHourglassSumArr); i++ {
	//	testHourglassSumArr[i].arr
	//}
}
