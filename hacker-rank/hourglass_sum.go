package main

import "math"

func hourglassSum(arr [][]int32) int32 {
	var max int32 = int32(math.Inf(-1))

	for i :=  0; i < 4; i++{
		for j := 0; j < 4; j++{
			top := arr[i][j] + arr[i][j+1] + arr[i][j+2]
			middle := arr[i+1][j+1]
			bottom := arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			subtotal := top + middle + bottom

			if subtotal > max {
				max = subtotal
			}
		}

	}
	return max
}