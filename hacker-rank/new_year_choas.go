package main

import "fmt"

//did not get to solve this problem
func minimumBribes (array []int32){
	//declare an empty slice
	var initialArray []int32

	//make an initial array with ascending order from 1 to n
	for i := 1; i <= len(array); i++{
		initialArray = append(initialArray,int32(i))
	}

	chaoticMap := make(map[int]bool)
	counter := 0
	tooChaotic := false

	for i := 0; i < len(array); i++{
		bribesCount := int(array[i]) - (i+1)
		if bribesCount > 0 && bribesCount <= 2{
			counter += bribesCount
			chaoticMap[counter] = tooChaotic
		} else if bribesCount > 0 && bribesCount > 2{
			tooChaotic = true
			counter += bribesCount
			chaoticMap[counter] = tooChaotic
		}
	}

	fmt.Println(chaoticMap)

	//if counter > 2 {
	//	fmt.Println(counter)
	//	fmt.Println("Too choatic")
	//}else{
	//	fmt.Println(counter)
	//}
}
