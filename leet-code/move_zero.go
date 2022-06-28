package main

func moveZero(nums []int) {
	position := 0
	nonZeroPtr := 0

	//move non-zero elements to the front of the array
	for nonZeroPtr < len(nums) {
		if nums[nonZeroPtr] != 0 {
			nums[position] = nums[nonZeroPtr]
			position++
		}
		nonZeroPtr++
	}

	//fill remaining array with 0
	for position < len(nums) {
		nums[position] = 0
		position++
	}
}
