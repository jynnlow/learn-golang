package main

func sortedSquaredArray(nums []int) []int {
	leftPtr := 0
	rightPtr := len(nums) - 1
	//using make function is to allocate given length array with all default values in the array
	//[0,0,0,0]
	sortedSquares := make([]int, len(nums))

	//append elements from the last index until index zero
	for i := len(nums) - 1; i >= 0; i-- {
		leftSquared := nums[leftPtr] * nums[leftPtr]
		rightSquared := nums[rightPtr] * nums[rightPtr]

		if leftSquared > rightSquared {
			//using append here will be adding element futher
			//[0,0,0,0,2]
			//we have to replace the element with index number but not appending
			sortedSquares[i] = leftSquared
			leftPtr++
		} else {
			sortedSquares[i] = rightSquared
			rightPtr--
		}
	}
	return sortedSquares
}
