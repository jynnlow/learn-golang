package main

func rotateArray(nums []int, k int) []int {
	for i := 0; i < k; i++ {
		current := nums[len(nums)-1]
		//j do not have to reach the first element (j>0)
		//when j hits the second element, the first element will be move to the second index
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = current
	}
	return nums
}
