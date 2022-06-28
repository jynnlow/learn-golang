package main

func twoSum(numbers []int, target int) []int {
	leftPtr := 0
	rightPtr := len(numbers) - 1

	for leftPtr < rightPtr {
		sum := numbers[leftPtr] + numbers[rightPtr]
		if target == sum {
			return []int{leftPtr + 1, rightPtr + 1}
		} else if target < sum {
			rightPtr--
		} else {
			leftPtr++
		}
	}
	return []int{}
}
