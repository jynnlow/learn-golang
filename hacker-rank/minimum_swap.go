package main

import "fmt"

//this solution is almost correct but it will have run time error when reach 10000 element in array
func minimumSwap (arr []int32) int32{
	var swap int32
	for i := 0; i < len(arr)-1; i++{
		for j := i+1;  j<len(arr); j++{
			if int(arr[j]) == i+1{
				temp := arr[j]
				arr[j] = arr[i]
				arr[i] = temp
				swap += 1
			}
		}
		fmt.Println(arr)
	}
	return swap
}

//this is final solution where a map acts as the variance array in each iteration
func minimumSwapUsingMap (arr []int32) int32{
	var swap int32
	arrayMap := make(map[int32]int)
	for i := 0; i < len(arr); i ++{
		arrayMap [arr[i]] = i
	}
	fmt.Println(arrayMap)

	for i := 0; i < len(arr)-1; i++{
		if int(arr[i]) != i+1 {
			incorrectValue := arr[i]
			arr[i] = int32(i + 1)
			arr[arrayMap[int32(i+1)]] = incorrectValue
			swap += 1

			arrayMap[incorrectValue] = arrayMap[int32(i+1)]
			arrayMap[int32(i+1)] = int(incorrectValue)
		}
	}
	return swap
}


