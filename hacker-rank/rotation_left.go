package main

import "fmt"

func main() {
	fmt.Println(rotLeftModulusSolution([]int32{1,2,3,4,5},2))
}

func rotLeft(array []int32, toLeft int32)[]int32{
	//declare an empty slice
	var rotatedArray []int32

	//declare a pointer j to select the element from the array given - start from toLeft till the end
	j := int(toLeft)

	//transfer the element in the array given into the slices
	for i := 0; i < len(array) - int(toLeft); i ++ {
		rotatedArray = append(rotatedArray,array[j])
		if j == len(array) - 1 {
			break
		}else{
			j++
		}
	}

	//change the pointer j to select the element from the beginning of the array until the toLeft - 1
	j = 0

	for i := len(rotatedArray); i < len(array); i ++{
		rotatedArray = append(rotatedArray,array[j])
		if j > int(toLeft){
			break
		}else{
			j++
		}
	}
	return rotatedArray
}

func rotLeftModulusSolution (array []int32, toLeft int32)[]int32 {
	//declare an empty slice
	rotatedArray := make([]int32, len(array))

	for i := 0; i < len(array); i++ {
		newLocation := (i + (len(array) - int(toLeft))) % len(array)
		rotatedArray[newLocation] = array[i]
	}
	return rotatedArray
}
