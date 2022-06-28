package main

func searchInputPosition(nums []int, target int) int {
	leftPtr := 0
	rightPtr := len(nums) - 1
	lastCheckingIdx := -1

	for leftPtr <= rightPtr {
		mid := (leftPtr + rightPtr) / 2
		lastCheckingIdx = mid
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			rightPtr = mid - 1
		} else {
			leftPtr = mid + 1
		}
	}

	if target < nums[lastCheckingIdx] {
		return lastCheckingIdx
	}
	return lastCheckingIdx + 1
}
